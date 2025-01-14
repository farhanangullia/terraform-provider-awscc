// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cleanrooms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cleanrooms_membership", membershipDataSource)
}

// membershipDataSource returns the Terraform awscc_cleanrooms_membership data source.
// This Terraform data source corresponds to the CloudFormation AWS::CleanRooms::Membership resource.
func membershipDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CollaborationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "type": "string"
		//	}
		"collaboration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CollaborationCreatorAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^\\d+$",
		//	  "type": "string"
		//	}
		"collaboration_creator_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CollaborationIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"collaboration_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MembershipIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"membership_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: QueryLogStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"query_log_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms membership.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An arbitrary set of tags (key-value pairs) for this cleanrooms membership.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CleanRooms::Membership",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CleanRooms::Membership").WithTerraformTypeName("awscc_cleanrooms_membership")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"collaboration_arn":                "CollaborationArn",
		"collaboration_creator_account_id": "CollaborationCreatorAccountId",
		"collaboration_identifier":         "CollaborationIdentifier",
		"key":                              "Key",
		"membership_identifier":            "MembershipIdentifier",
		"query_log_status":                 "QueryLogStatus",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
