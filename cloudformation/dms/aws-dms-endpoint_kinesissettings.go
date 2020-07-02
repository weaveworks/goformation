package dms

import (
	"github.com/weaveworks/goformation/v4/cloudformation/types"

	"github.com/weaveworks/goformation/v4/cloudformation/policies"
)

// Endpoint_KinesisSettings AWS CloudFormation Resource (AWS::DMS::Endpoint.KinesisSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html
type Endpoint_KinesisSettings struct {

	// MessageFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-messageformat
	MessageFormat *types.Value `json:"MessageFormat,omitempty"`

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-serviceaccessrolearn
	ServiceAccessRoleArn *types.Value `json:"ServiceAccessRoleArn,omitempty"`

	// StreamArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-streamarn
	StreamArn *types.Value `json:"StreamArn,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Endpoint_KinesisSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.KinesisSettings"
}
