// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codedeploy_deployment_config", deploymentConfigDataSource)
}

// deploymentConfigDataSource returns the Terraform awscc_codedeploy_deployment_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeDeploy::DeploymentConfig resource.
func deploymentConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"compute_platform": {
			// Property: ComputePlatform
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The destination platform type for the deployment (Lambda, Server, or ECS).",
			//	  "type": "string"
			//	}
			Description: "The destination platform type for the deployment (Lambda, Server, or ECS).",
			Type:        types.StringType,
			Computed:    true,
		},
		"deployment_config_name": {
			// Property: DeploymentConfigName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
			//	  "type": "string"
			//	}
			Description: "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"minimum_healthy_hosts": {
			// Property: MinimumHealthyHosts
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
			//	  "properties": {
			//	    "Type": {
			//	      "type": "string"
			//	    },
			//	    "Value": {
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "required": [
			//	    "Type",
			//	    "Value"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.Int64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"traffic_routing_config": {
			// Property: TrafficRoutingConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The configuration that specifies how the deployment traffic is routed.",
			//	  "properties": {
			//	    "TimeBasedCanary": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CanaryInterval": {
			//	          "type": "integer"
			//	        },
			//	        "CanaryPercentage": {
			//	          "type": "integer"
			//	        }
			//	      },
			//	      "required": [
			//	        "CanaryPercentage",
			//	        "CanaryInterval"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "TimeBasedLinear": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "LinearInterval": {
			//	          "type": "integer"
			//	        },
			//	        "LinearPercentage": {
			//	          "type": "integer"
			//	        }
			//	      },
			//	      "required": [
			//	        "LinearInterval",
			//	        "LinearPercentage"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "Type": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Type"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The configuration that specifies how the deployment traffic is routed.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"time_based_canary": {
						// Property: TimeBasedCanary
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"canary_interval": {
									// Property: CanaryInterval
									Type:     types.Int64Type,
									Computed: true,
								},
								"canary_percentage": {
									// Property: CanaryPercentage
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"time_based_linear": {
						// Property: TimeBasedLinear
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"linear_interval": {
									// Property: LinearInterval
									Type:     types.Int64Type,
									Computed: true,
								},
								"linear_percentage": {
									// Property: LinearPercentage
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
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
		Description: "Data Source schema for AWS::CodeDeploy::DeploymentConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeDeploy::DeploymentConfig").WithTerraformTypeName("awscc_codedeploy_deployment_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"canary_interval":        "CanaryInterval",
		"canary_percentage":      "CanaryPercentage",
		"compute_platform":       "ComputePlatform",
		"deployment_config_name": "DeploymentConfigName",
		"linear_interval":        "LinearInterval",
		"linear_percentage":      "LinearPercentage",
		"minimum_healthy_hosts":  "MinimumHealthyHosts",
		"time_based_canary":      "TimeBasedCanary",
		"time_based_linear":      "TimeBasedLinear",
		"traffic_routing_config": "TrafficRoutingConfig",
		"type":                   "Type",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}