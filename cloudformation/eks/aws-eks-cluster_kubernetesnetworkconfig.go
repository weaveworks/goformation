package eks

import (
	"github.com/weaveworks/goformation/v4/cloudformation/types"

	"github.com/weaveworks/goformation/v4/cloudformation/policies"
)

// Cluster_KubernetesNetworkConfig AWS CloudFormation Resource (AWS::EKS::Cluster.KubernetesNetworkConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubernetesnetworkconfig.html
type Cluster_KubernetesNetworkConfig struct {

	// ServiceIpv4Cidr AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubernetesnetworkconfig.html#cfn-eks-cluster-kubernetesnetworkconfig-serviceipv4cidr
	ServiceIpv4Cidr *types.Value `json:"ServiceIpv4Cidr,omitempty"`

	// ServiceIpv6Cidr represents the CIDR block to assign the Kubernetes Service IP addresses from
	// Expects a string
	ServiceIpv6Cidr *types.Value `json:"ServiceIpv6Cidr,omitempty"`

	// IpFamily can only be "ipv4" or "ipv6"
	// IPv6 is only supported on clusters with k8s version 1.21 or higher
	IpFamily *types.Value `json:"IpFamily,omitempty"`

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
func (r *Cluster_KubernetesNetworkConfig) AWSCloudFormationType() string {
	return "AWS::EKS::Cluster.KubernetesNetworkConfig"
}
