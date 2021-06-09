package eks

import "github.com/weaveworks/goformation/v4/cloudformation/types"

// Nodegroup_UpdateConfig AWS CloudFormation Resource (AWS::EKS::Nodegroup.UpdateConfig)
// TODO @nikimanoledaki: "See" with link to AWS CF docs
type Nodegroup_UpdateConfig struct {

	// MaxUnavailable AWS CloudFormation Property
	// Required: false
	// TODO: add link to docs
	MaxUnavailable *types.Value `json:"MaxUnavailable,omitempty"`

	// MaxUnavailablePercentage AWS CloudFormation Property
	// Required: false
	// TODO: add link to docs
	MaxUnavailablePercentage *types.Value `json:"MaxUnavailablePercentage,omitempty"`
}
