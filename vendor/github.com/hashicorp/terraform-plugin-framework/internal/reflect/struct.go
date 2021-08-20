package reflect

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/internal/diagnostics"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Struct builds a new struct using the data in `object`, as long as `object`
// is a `tftypes.Object`. It will take the struct type from `target`, which
// must be a struct type.
//
// The properties on `target` must be tagged with a "tfsdk" label containing
// the field name to map to that property. Every property must be tagged, and
// every property must be present in the type of `object`, and all the
// attributes in the type of `object` must have a corresponding property.
// Properties that don't map to object attributes must have a `tfsdk:"-"` tag,
// explicitly defining them as not part of the object. This is to catch typos
// and other mistakes early.
//
// Struct is meant to be called from Into, not directly.
func Struct(ctx context.Context, typ attr.Type, object tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	// this only works with object values, so make sure that constraint is
	// met
	if target.Kind() != reflect.Struct {
		err := fmt.Errorf("expected a struct type, got %s", target.Type())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	if !object.Type().Is(tftypes.Object{}) {
		err := fmt.Errorf("cannot reflect %s into a struct, must be an object", object.Type().String())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	attrsType, ok := typ.(attr.TypeWithAttributeTypes)
	if !ok {
		err := fmt.Errorf("cannot reflect object using type information provided by %T, %T must be an attr.TypeWithAttributeTypes", typ, typ)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}

	// collect a map of fields that are in the object passed in
	var objectFields map[string]tftypes.Value
	err := object.As(&objectFields)
	if err != nil {
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}

	// collect a map of fields that are defined in the tags of the struct
	// passed in
	targetFields, err := getStructTags(ctx, target, path)
	if err != nil {
		err = fmt.Errorf("error retrieving field names from struct tags: %w", err)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}

	// we require an exact, 1:1 match of these fields to avoid typos
	// leading to surprises, so let's ensure they have the exact same
	// fields defined
	var objectMissing, targetMissing []string
	for field := range targetFields {
		if _, ok := objectFields[field]; !ok {
			objectMissing = append(objectMissing, field)
		}
	}
	for field := range objectFields {
		if _, ok := targetFields[field]; !ok {
			targetMissing = append(targetMissing, field)
		}
	}
	if len(objectMissing) > 0 || len(targetMissing) > 0 {
		var missing []string
		if len(objectMissing) > 0 {
			missing = append(missing, fmt.Sprintf("Struct defines fields not found in object: %s.", commaSeparatedString(objectMissing)))
		}
		if len(targetMissing) > 0 {
			missing = append(missing, fmt.Sprintf("Object defines fields not found in struct: %s.", commaSeparatedString(targetMissing)))
		}
		err := fmt.Errorf("mismatch between struct and object: %s", strings.Join(missing, " "))
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}

	attrTypes := attrsType.AttributeTypes()

	// now that we know they match perfectly, fill the struct with the
	// values in the object
	result := reflect.New(target.Type()).Elem()
	for field, structFieldPos := range targetFields {
		attrType, ok := attrTypes[field]
		if !ok {
			err := fmt.Errorf("could not find type information for attribute in supplied attr.Type %T", typ)
			return target, append(diags, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Value Conversion Error",
				Detail:    "An unexpected error was encountered trying to convert to struct. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
				Attribute: path.WithAttributeName(field),
			})
		}
		structField := result.Field(structFieldPos)
		fieldVal, fieldValDiags := BuildValue(ctx, attrType, objectFields[field], structField, opts, path.WithAttributeName(field))
		diags = append(diags, fieldValDiags...)

		if diagnostics.DiagsHasErrors(diags) {
			return target, diags
		}
		structField.Set(fieldVal)
	}
	return result, diags
}

// FromStruct builds an attr.Value as produced by `typ` from the data in `val`.
// `val` must be a struct type, and must have all its properties tagged and be
// a 1:1 match with the attributes reported by `typ`. FromStruct will recurse
// into FromValue for each attribute, using the type of the attribute as
// reported by `typ`.
//
// It is meant to be called through OutOf, not directly.
func FromStruct(ctx context.Context, typ attr.TypeWithAttributeTypes, val reflect.Value, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	objTypes := map[string]tftypes.Type{}
	objValues := map[string]tftypes.Value{}

	// collect a map of fields that are defined in the tags of the struct
	// passed in
	targetFields, err := getStructTags(ctx, val, path)
	if err != nil {
		err = fmt.Errorf("error retrieving field names from struct tags: %w", err)
		return nil, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert from struct value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}

	attrTypes := typ.AttributeTypes()
	for name, fieldNo := range targetFields {
		path := path.WithAttributeName(name)
		fieldValue := val.Field(fieldNo)

		attrVal, attrValDiags := FromValue(ctx, attrTypes[name], fieldValue.Interface(), path)
		diags = append(diags, attrValDiags...)

		if diagnostics.DiagsHasErrors(diags) {
			return nil, diags
		}

		attrType, ok := attrTypes[name]
		if !ok || attrType == nil {
			err := fmt.Errorf("couldn't find type information for attribute in supplied attr.Type %T", typ)
			return nil, append(diags, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Value Conversion Error",
				Detail:    "An unexpected error was encountered trying to convert from struct value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
				Attribute: path,
			})
		}

		objTypes[name] = attrType.TerraformType(ctx)

		tfVal, err := attrVal.ToTerraformValue(ctx)
		if err != nil {
			return nil, append(diags, toTerraformValueErrorDiag(err, path))
		}
		err = tftypes.ValidateValue(objTypes[name], tfVal)
		if err != nil {
			return nil, append(diags, validateValueErrorDiag(err, path))
		}

		tfObjVal := tftypes.NewValue(objTypes[name], tfVal)

		if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
			diags = append(diags, typeWithValidate.Validate(ctx, tfObjVal)...)

			if diagnostics.DiagsHasErrors(diags) {
				return nil, diags
			}
		}

		objValues[name] = tfObjVal
	}

	tfVal := tftypes.NewValue(tftypes.Object{
		AttributeTypes: objTypes,
	}, objValues)

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

		if diagnostics.DiagsHasErrors(diags) {
			return nil, diags
		}
	}

	retType := typ.WithAttributeTypes(attrTypes)
	ret, err := retType.ValueFromTerraform(ctx, tfVal)
	if err != nil {
		return nil, append(diags, valueFromTerraformErrorDiag(err, path))
	}

	return ret, diags
}
