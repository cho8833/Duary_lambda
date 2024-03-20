// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status for contributor insights for a specific table or index.
// CloudWatch Contributor Insights for DynamoDB graphs display the partition key
// and (if applicable) sort key of frequently accessed items and frequently
// throttled items in plaintext. If you require the use of Amazon Web Services Key
// Management Service (KMS) to encrypt this table’s partition key and sort key data
// with an Amazon Web Services managed key or customer managed key, you should not
// enable CloudWatch Contributor Insights for DynamoDB for this table.
func (c *Client) UpdateContributorInsights(ctx context.Context, params *UpdateContributorInsightsInput, optFns ...func(*Options)) (*UpdateContributorInsightsOutput, error) {
	if params == nil {
		params = &UpdateContributorInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContributorInsights", params, optFns, c.addOperationUpdateContributorInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContributorInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContributorInsightsInput struct {

	// Represents the contributor insights action.
	//
	// This member is required.
	ContributorInsightsAction types.ContributorInsightsAction

	// The name of the table.
	//
	// This member is required.
	TableName *string

	// The global secondary index name, if applicable.
	IndexName *string

	noSmithyDocumentSerde
}

type UpdateContributorInsightsOutput struct {

	// The status of contributor insights
	ContributorInsightsStatus types.ContributorInsightsStatus

	// The name of the global secondary index, if applicable.
	IndexName *string

	// The name of the table.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContributorInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateContributorInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateContributorInsights{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContributorInsights"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateContributorInsightsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContributorInsights(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateContributorInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContributorInsights",
	}
}
