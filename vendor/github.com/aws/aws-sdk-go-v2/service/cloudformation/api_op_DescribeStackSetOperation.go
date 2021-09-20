// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of the specified stack set operation.
func (c *Client) DescribeStackSetOperation(ctx context.Context, params *DescribeStackSetOperationInput, optFns ...func(*Options)) (*DescribeStackSetOperationOutput, error) {
	if params == nil {
		params = &DescribeStackSetOperationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackSetOperation", params, optFns, c.addOperationDescribeStackSetOperationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackSetOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackSetOperationInput struct {

	// The unique ID of the stack set operation.
	//
	// This member is required.
	OperationId *string

	// The name or the unique stack ID of the stack set for the stack operation.
	//
	// This member is required.
	StackSetName *string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * If you are signed in to the
	// management account, specify SELF.
	//
	// * If you are signed in to a delegated
	// administrator account, specify DELEGATED_ADMIN. Your AWS account must be
	// registered as a delegated administrator in the management account. For more
	// information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the AWS CloudFormation User Guide.
	CallAs types.CallAs
}

type DescribeStackSetOperationOutput struct {

	// The specified stack set operation.
	StackSetOperation *types.StackSetOperation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeStackSetOperationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackSetOperation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackSetOperation{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeStackSetOperationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackSetOperation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeStackSetOperation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackSetOperation",
	}
}
