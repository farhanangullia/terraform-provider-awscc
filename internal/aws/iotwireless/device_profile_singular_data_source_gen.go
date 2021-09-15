// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotwireless_device_profile", deviceProfileDataSourceType)
}

// deviceProfileDataSourceType returns the Terraform awscc_iotwireless_device_profile data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTWireless::DeviceProfile resource type.
func deviceProfileDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Service profile Arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "Service profile Arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Service profile Id. Returned after successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Service profile Id. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ClassBTimeout": {
			//       "maximum": 1000,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "ClassCTimeout": {
			//       "maximum": 1000,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MacVersion": {
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "MaxDutyCycle": {
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MaxEirp": {
			//       "maximum": 15,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "PingSlotDr": {
			//       "maximum": 15,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "PingSlotFreq": {
			//       "maximum": 16700000,
			//       "minimum": 1000000,
			//       "type": "integer"
			//     },
			//     "PingSlotPeriod": {
			//       "maximum": 4096,
			//       "minimum": 128,
			//       "type": "integer"
			//     },
			//     "RegParamsRevision": {
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "RfRegion": {
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "Supports32BitFCnt": {
			//       "type": "boolean"
			//     },
			//     "SupportsClassB": {
			//       "type": "boolean"
			//     },
			//     "SupportsClassC": {
			//       "type": "boolean"
			//     },
			//     "SupportsJoin": {
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"class_b_timeout": {
						// Property: ClassBTimeout
						Type:     types.NumberType,
						Computed: true,
					},
					"class_c_timeout": {
						// Property: ClassCTimeout
						Type:     types.NumberType,
						Computed: true,
					},
					"mac_version": {
						// Property: MacVersion
						Type:     types.StringType,
						Computed: true,
					},
					"max_duty_cycle": {
						// Property: MaxDutyCycle
						Type:     types.NumberType,
						Computed: true,
					},
					"max_eirp": {
						// Property: MaxEirp
						Type:     types.NumberType,
						Computed: true,
					},
					"ping_slot_dr": {
						// Property: PingSlotDr
						Type:     types.NumberType,
						Computed: true,
					},
					"ping_slot_freq": {
						// Property: PingSlotFreq
						Type:     types.NumberType,
						Computed: true,
					},
					"ping_slot_period": {
						// Property: PingSlotPeriod
						Type:     types.NumberType,
						Computed: true,
					},
					"reg_params_revision": {
						// Property: RegParamsRevision
						Type:     types.StringType,
						Computed: true,
					},
					"rf_region": {
						// Property: RfRegion
						Type:     types.StringType,
						Computed: true,
					},
					"supports_32_bit_f_cnt": {
						// Property: Supports32BitFCnt
						Type:     types.BoolType,
						Computed: true,
					},
					"supports_class_b": {
						// Property: SupportsClassB
						Type:     types.BoolType,
						Computed: true,
					},
					"supports_class_c": {
						// Property: SupportsClassC
						Type:     types.BoolType,
						Computed: true,
					},
					"supports_join": {
						// Property: SupportsJoin
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of service profile",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Name of service profile",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the device profile.",
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
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the device profile.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTWireless::DeviceProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::DeviceProfile").WithTerraformTypeName("awscc_iotwireless_device_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"class_b_timeout":       "ClassBTimeout",
		"class_c_timeout":       "ClassCTimeout",
		"id":                    "Id",
		"key":                   "Key",
		"lo_ra_wan":             "LoRaWAN",
		"mac_version":           "MacVersion",
		"max_duty_cycle":        "MaxDutyCycle",
		"max_eirp":              "MaxEirp",
		"name":                  "Name",
		"ping_slot_dr":          "PingSlotDr",
		"ping_slot_freq":        "PingSlotFreq",
		"ping_slot_period":      "PingSlotPeriod",
		"reg_params_revision":   "RegParamsRevision",
		"rf_region":             "RfRegion",
		"supports_32_bit_f_cnt": "Supports32BitFCnt",
		"supports_class_b":      "SupportsClassB",
		"supports_class_c":      "SupportsClassC",
		"supports_join":         "SupportsJoin",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}