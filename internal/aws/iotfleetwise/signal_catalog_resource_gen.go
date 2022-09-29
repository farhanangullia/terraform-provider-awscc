// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotfleetwise_signal_catalog", signalCatalogResourceType)
}

// signalCatalogResourceType returns the Terraform awscc_iotfleetwise_signal_catalog resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTFleetWise::SignalCatalog resource type.
func signalCatalogResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "format": "date-time",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_modification_time": {
			// Property: LastModificationTime
			// CloudFormation resource type schema:
			// {
			//   "format": "date-time",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z\\d\\-_:]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z\\d\\-_:]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"node_counts": {
			// Property: NodeCounts
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "TotalActuators": {
			//       "type": "number"
			//     },
			//     "TotalAttributes": {
			//       "type": "number"
			//     },
			//     "TotalBranches": {
			//       "type": "number"
			//     },
			//     "TotalNodes": {
			//       "type": "number"
			//     },
			//     "TotalSensors": {
			//       "type": "number"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"total_actuators": {
						// Property: TotalActuators
						Type:     types.Float64Type,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"total_attributes": {
						// Property: TotalAttributes
						Type:     types.Float64Type,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"total_branches": {
						// Property: TotalBranches
						Type:     types.Float64Type,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"total_nodes": {
						// Property: TotalNodes
						Type:     types.Float64Type,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"total_sensors": {
						// Property: TotalSensors
						Type:     types.Float64Type,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"nodes": {
			// Property: Nodes
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "properties": {
			//       "Actuator": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AllowedValues": {
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "AssignedValue": {
			//             "type": "string"
			//           },
			//           "DataType": {
			//             "enum": [
			//               "INT8",
			//               "UINT8",
			//               "INT16",
			//               "UINT16",
			//               "INT32",
			//               "UINT32",
			//               "INT64",
			//               "UINT64",
			//               "BOOLEAN",
			//               "FLOAT",
			//               "DOUBLE",
			//               "STRING",
			//               "UNIX_TIMESTAMP",
			//               "INT8_ARRAY",
			//               "UINT8_ARRAY",
			//               "INT16_ARRAY",
			//               "UINT16_ARRAY",
			//               "INT32_ARRAY",
			//               "UINT32_ARRAY",
			//               "INT64_ARRAY",
			//               "UINT64_ARRAY",
			//               "BOOLEAN_ARRAY",
			//               "FLOAT_ARRAY",
			//               "DOUBLE_ARRAY",
			//               "STRING_ARRAY",
			//               "UNIX_TIMESTAMP_ARRAY",
			//               "UNKNOWN"
			//             ],
			//             "type": "string"
			//           },
			//           "Description": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "FullyQualifiedName": {
			//             "type": "string"
			//           },
			//           "Max": {
			//             "type": "number"
			//           },
			//           "Min": {
			//             "type": "number"
			//           },
			//           "Unit": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "DataType",
			//           "FullyQualifiedName"
			//         ],
			//         "type": "object"
			//       },
			//       "Attribute": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AllowedValues": {
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "AssignedValue": {
			//             "type": "string"
			//           },
			//           "DataType": {
			//             "enum": [
			//               "INT8",
			//               "UINT8",
			//               "INT16",
			//               "UINT16",
			//               "INT32",
			//               "UINT32",
			//               "INT64",
			//               "UINT64",
			//               "BOOLEAN",
			//               "FLOAT",
			//               "DOUBLE",
			//               "STRING",
			//               "UNIX_TIMESTAMP",
			//               "INT8_ARRAY",
			//               "UINT8_ARRAY",
			//               "INT16_ARRAY",
			//               "UINT16_ARRAY",
			//               "INT32_ARRAY",
			//               "UINT32_ARRAY",
			//               "INT64_ARRAY",
			//               "UINT64_ARRAY",
			//               "BOOLEAN_ARRAY",
			//               "FLOAT_ARRAY",
			//               "DOUBLE_ARRAY",
			//               "STRING_ARRAY",
			//               "UNIX_TIMESTAMP_ARRAY",
			//               "UNKNOWN"
			//             ],
			//             "type": "string"
			//           },
			//           "DefaultValue": {
			//             "type": "string"
			//           },
			//           "Description": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "FullyQualifiedName": {
			//             "type": "string"
			//           },
			//           "Max": {
			//             "type": "number"
			//           },
			//           "Min": {
			//             "type": "number"
			//           },
			//           "Unit": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "DataType",
			//           "FullyQualifiedName"
			//         ],
			//         "type": "object"
			//       },
			//       "Branch": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Description": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "FullyQualifiedName": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "FullyQualifiedName"
			//         ],
			//         "type": "object"
			//       },
			//       "Sensor": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AllowedValues": {
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "DataType": {
			//             "enum": [
			//               "INT8",
			//               "UINT8",
			//               "INT16",
			//               "UINT16",
			//               "INT32",
			//               "UINT32",
			//               "INT64",
			//               "UINT64",
			//               "BOOLEAN",
			//               "FLOAT",
			//               "DOUBLE",
			//               "STRING",
			//               "UNIX_TIMESTAMP",
			//               "INT8_ARRAY",
			//               "UINT8_ARRAY",
			//               "INT16_ARRAY",
			//               "UINT16_ARRAY",
			//               "INT32_ARRAY",
			//               "UINT32_ARRAY",
			//               "INT64_ARRAY",
			//               "UINT64_ARRAY",
			//               "BOOLEAN_ARRAY",
			//               "FLOAT_ARRAY",
			//               "DOUBLE_ARRAY",
			//               "STRING_ARRAY",
			//               "UNIX_TIMESTAMP_ARRAY",
			//               "UNKNOWN"
			//             ],
			//             "type": "string"
			//           },
			//           "Description": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "FullyQualifiedName": {
			//             "type": "string"
			//           },
			//           "Max": {
			//             "type": "number"
			//           },
			//           "Min": {
			//             "type": "number"
			//           },
			//           "Unit": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "DataType",
			//           "FullyQualifiedName"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 500,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"actuator": {
						// Property: Actuator
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"allowed_values": {
									// Property: AllowedValues
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"assigned_value": {
									// Property: AssignedValue
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"data_type": {
									// Property: DataType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"INT8",
											"UINT8",
											"INT16",
											"UINT16",
											"INT32",
											"UINT32",
											"INT64",
											"UINT64",
											"BOOLEAN",
											"FLOAT",
											"DOUBLE",
											"STRING",
											"UNIX_TIMESTAMP",
											"INT8_ARRAY",
											"UINT8_ARRAY",
											"INT16_ARRAY",
											"UINT16_ARRAY",
											"INT32_ARRAY",
											"UINT32_ARRAY",
											"INT64_ARRAY",
											"UINT64_ARRAY",
											"BOOLEAN_ARRAY",
											"FLOAT_ARRAY",
											"DOUBLE_ARRAY",
											"STRING_ARRAY",
											"UNIX_TIMESTAMP_ARRAY",
											"UNKNOWN",
										}),
									},
								},
								"description": {
									// Property: Description
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"fully_qualified_name": {
									// Property: FullyQualifiedName
									Type:     types.StringType,
									Required: true,
								},
								"max": {
									// Property: Max
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"min": {
									// Property: Min
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"unit": {
									// Property: Unit
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"attribute": {
						// Property: Attribute
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"allowed_values": {
									// Property: AllowedValues
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"assigned_value": {
									// Property: AssignedValue
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"data_type": {
									// Property: DataType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"INT8",
											"UINT8",
											"INT16",
											"UINT16",
											"INT32",
											"UINT32",
											"INT64",
											"UINT64",
											"BOOLEAN",
											"FLOAT",
											"DOUBLE",
											"STRING",
											"UNIX_TIMESTAMP",
											"INT8_ARRAY",
											"UINT8_ARRAY",
											"INT16_ARRAY",
											"UINT16_ARRAY",
											"INT32_ARRAY",
											"UINT32_ARRAY",
											"INT64_ARRAY",
											"UINT64_ARRAY",
											"BOOLEAN_ARRAY",
											"FLOAT_ARRAY",
											"DOUBLE_ARRAY",
											"STRING_ARRAY",
											"UNIX_TIMESTAMP_ARRAY",
											"UNKNOWN",
										}),
									},
								},
								"default_value": {
									// Property: DefaultValue
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"description": {
									// Property: Description
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"fully_qualified_name": {
									// Property: FullyQualifiedName
									Type:     types.StringType,
									Required: true,
								},
								"max": {
									// Property: Max
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"min": {
									// Property: Min
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"unit": {
									// Property: Unit
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"branch": {
						// Property: Branch
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"description": {
									// Property: Description
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"fully_qualified_name": {
									// Property: FullyQualifiedName
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"sensor": {
						// Property: Sensor
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"allowed_values": {
									// Property: AllowedValues
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"data_type": {
									// Property: DataType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"INT8",
											"UINT8",
											"INT16",
											"UINT16",
											"INT32",
											"UINT32",
											"INT64",
											"UINT64",
											"BOOLEAN",
											"FLOAT",
											"DOUBLE",
											"STRING",
											"UNIX_TIMESTAMP",
											"INT8_ARRAY",
											"UINT8_ARRAY",
											"INT16_ARRAY",
											"UINT16_ARRAY",
											"INT32_ARRAY",
											"UINT32_ARRAY",
											"INT64_ARRAY",
											"UINT64_ARRAY",
											"BOOLEAN_ARRAY",
											"FLOAT_ARRAY",
											"DOUBLE_ARRAY",
											"STRING_ARRAY",
											"UNIX_TIMESTAMP_ARRAY",
											"UNKNOWN",
										}),
									},
								},
								"description": {
									// Property: Description
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"fully_qualified_name": {
									// Property: FullyQualifiedName
									Type:     types.StringType,
									Required: true,
								},
								"max": {
									// Property: Max
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"min": {
									// Property: Min
									Type:     types.Float64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"unit": {
									// Property: Unit
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 500),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::IoTFleetWise::SignalCatalog Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::SignalCatalog").WithTerraformTypeName("awscc_iotfleetwise_signal_catalog")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actuator":               "Actuator",
		"allowed_values":         "AllowedValues",
		"arn":                    "Arn",
		"assigned_value":         "AssignedValue",
		"attribute":              "Attribute",
		"branch":                 "Branch",
		"creation_time":          "CreationTime",
		"data_type":              "DataType",
		"default_value":          "DefaultValue",
		"description":            "Description",
		"fully_qualified_name":   "FullyQualifiedName",
		"key":                    "Key",
		"last_modification_time": "LastModificationTime",
		"max":                    "Max",
		"min":                    "Min",
		"name":                   "Name",
		"node_counts":            "NodeCounts",
		"nodes":                  "Nodes",
		"sensor":                 "Sensor",
		"tags":                   "Tags",
		"total_actuators":        "TotalActuators",
		"total_attributes":       "TotalAttributes",
		"total_branches":         "TotalBranches",
		"total_nodes":            "TotalNodes",
		"total_sensors":          "TotalSensors",
		"unit":                   "Unit",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
