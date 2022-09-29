// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotwireless_device_profile", deviceProfileResourceType)
}

// deviceProfileResourceType returns the Terraform awscc_iotwireless_device_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::DeviceProfile resource type.
func deviceProfileResourceType(ctx context.Context) (provider.ResourceType, error) {
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation",
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
			//     "FactoryPresetFreqsList": {
			//       "items": {
			//         "maximum": 16700000,
			//         "minimum": 1000000,
			//         "type": "integer"
			//       },
			//       "maxItems": 20,
			//       "type": "array"
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
			//     "RxDataRate2": {
			//       "maximum": 15,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "RxDelay1": {
			//       "maximum": 15,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "RxDrOffset1": {
			//       "maximum": 7,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "RxFreq2": {
			//       "maximum": 16700000,
			//       "minimum": 1000000,
			//       "type": "integer"
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
			Description: "LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"class_b_timeout": {
						// Property: ClassBTimeout
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 1000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"class_c_timeout": {
						// Property: ClassCTimeout
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 1000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"factory_preset_freqs_list": {
						// Property: FactoryPresetFreqsList
						Type:     types.ListType{ElemType: types.Int64Type},
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(20),
							validate.ArrayForEach(validate.IntBetween(1000000, 16700000)),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"mac_version": {
						// Property: MacVersion
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(64),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_duty_cycle": {
						// Property: MaxDutyCycle
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 100),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_eirp": {
						// Property: MaxEirp
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 15),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ping_slot_dr": {
						// Property: PingSlotDr
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 15),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ping_slot_freq": {
						// Property: PingSlotFreq
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1000000, 16700000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ping_slot_period": {
						// Property: PingSlotPeriod
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(128, 4096),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"reg_params_revision": {
						// Property: RegParamsRevision
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(64),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rf_region": {
						// Property: RfRegion
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(64),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rx_data_rate_2": {
						// Property: RxDataRate2
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 15),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rx_delay_1": {
						// Property: RxDelay1
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 15),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rx_dr_offset_1": {
						// Property: RxDrOffset1
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 7),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rx_freq_2": {
						// Property: RxFreq2
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1000000, 16700000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"supports_32_bit_f_cnt": {
						// Property: Supports32BitFCnt
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"supports_class_b": {
						// Property: SupportsClassB
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"supports_class_c": {
						// Property: SupportsClassC
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"supports_join": {
						// Property: SupportsJoin
						Type:     types.BoolType,
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
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the device profile.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Device Profile's resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::DeviceProfile").WithTerraformTypeName("awscc_iotwireless_device_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"class_b_timeout":           "ClassBTimeout",
		"class_c_timeout":           "ClassCTimeout",
		"factory_preset_freqs_list": "FactoryPresetFreqsList",
		"id":                        "Id",
		"key":                       "Key",
		"lo_ra_wan":                 "LoRaWAN",
		"mac_version":               "MacVersion",
		"max_duty_cycle":            "MaxDutyCycle",
		"max_eirp":                  "MaxEirp",
		"name":                      "Name",
		"ping_slot_dr":              "PingSlotDr",
		"ping_slot_freq":            "PingSlotFreq",
		"ping_slot_period":          "PingSlotPeriod",
		"reg_params_revision":       "RegParamsRevision",
		"rf_region":                 "RfRegion",
		"rx_data_rate_2":            "RxDataRate2",
		"rx_delay_1":                "RxDelay1",
		"rx_dr_offset_1":            "RxDrOffset1",
		"rx_freq_2":                 "RxFreq2",
		"supports_32_bit_f_cnt":     "Supports32BitFCnt",
		"supports_class_b":          "SupportsClassB",
		"supports_class_c":          "SupportsClassC",
		"supports_join":             "SupportsJoin",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
