// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudformation_module_default_version", moduleDefaultVersionResourceType)
}

// moduleDefaultVersionResourceType returns the Terraform awscc_cloudformation_module_default_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFormation::ModuleDefaultVersion resource type.
func moduleDefaultVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the module version to set as the default version.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the module version to set as the default version.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"module_name": {
			// Property: ModuleName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a module existing in the registry.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of a module existing in the registry.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// ModuleName is a write-only property.
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of an existing version of the named module to set as the default.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of an existing version of the named module to set as the default.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// VersionId is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "A module that has been registered in the CloudFormation registry as the default version",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ModuleDefaultVersion").WithTerraformTypeName("awscc_cloudformation_module_default_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"module_name": "ModuleName",
		"version_id":  "VersionId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ModuleName",
		"/properties/VersionId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"arn",
		),
		validate.Required(
			"module_name",
			"version_id",
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}