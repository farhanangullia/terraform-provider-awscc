// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iot_ca_certificate", cACertificateDataSourceType)
}

// cACertificateDataSourceType returns the Terraform awscc_iot_ca_certificate data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoT::CACertificate resource type.
func cACertificateDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"auto_registration_status": {
			// Property: AutoRegistrationStatus
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENABLE",
			//     "DISABLE"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ca_certificate_pem": {
			// Property: CACertificatePem
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "pattern": "[\\s\\S]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_mode": {
			// Property: CertificateMode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "SNI_ONLY"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"registration_config": {
			// Property: RegistrationConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "RoleArn": {
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "arn:(aws[a-zA-Z-]*)?:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
			//       "type": "string"
			//     },
			//     "TemplateBody": {
			//       "maxLength": 10240,
			//       "minLength": 0,
			//       "pattern": "[\\s\\S]*",
			//       "type": "string"
			//     },
			//     "TemplateName": {
			//       "maxLength": 36,
			//       "minLength": 1,
			//       "pattern": "^[0-9A-Za-z_-]+$",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"role_arn": {
						// Property: RoleArn
						Type:     types.StringType,
						Computed: true,
					},
					"template_body": {
						// Property: TemplateBody
						Type:     types.StringType,
						Computed: true,
					},
					"template_name": {
						// Property: TemplateName
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"remove_auto_registration": {
			// Property: RemoveAutoRegistration
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"verification_certificate_pem": {
			// Property: VerificationCertificatePem
			// CloudFormation resource type schema:
			// {
			//   "description": "The private key verification certificate.",
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "pattern": "[\\s\\S]*",
			//   "type": "string"
			// }
			Description: "The private key verification certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoT::CACertificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::CACertificate").WithTerraformTypeName("awscc_iot_ca_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"auto_registration_status":     "AutoRegistrationStatus",
		"ca_certificate_pem":           "CACertificatePem",
		"certificate_mode":             "CertificateMode",
		"id":                           "Id",
		"key":                          "Key",
		"registration_config":          "RegistrationConfig",
		"remove_auto_registration":     "RemoveAutoRegistration",
		"role_arn":                     "RoleArn",
		"status":                       "Status",
		"tags":                         "Tags",
		"template_body":                "TemplateBody",
		"template_name":                "TemplateName",
		"value":                        "Value",
		"verification_certificate_pem": "VerificationCertificatePem",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
