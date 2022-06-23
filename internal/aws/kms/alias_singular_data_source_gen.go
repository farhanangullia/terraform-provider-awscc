// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kms_alias", aliasDataSourceType)
}

// aliasDataSourceType returns the Terraform awscc_kms_alias data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::KMS::Alias resource type.
func aliasDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias_name": {
			// Property: AliasName
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^(alias/)[a-zA-Z0-9:/_-]+$",
			//   "type": "string"
			// }
			Description: "Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_key_id": {
			// Property: TargetKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
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
		Description: "Data Source schema for AWS::KMS::Alias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::Alias").WithTerraformTypeName("awscc_kms_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_name":    "AliasName",
		"target_key_id": "TargetKeyId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
