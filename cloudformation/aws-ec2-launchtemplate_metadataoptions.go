package cloudformation

// AWSEC2LaunchTemplate_MetadataOptions AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.MetadataOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html
type AWSEC2LaunchTemplate_MetadataOptions struct {

	// HttpEndpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpendpoint
	HttpEndpoint *Value `json:"HttpEndpoint,omitempty"`

	// HttpPutResponseHopLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpputresponsehoplimit
	HttpPutResponseHopLimit *Value `json:"HttpPutResponseHopLimit,omitempty"`

	// HttpTokens AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens
	HttpTokens *Value `json:"HttpTokens,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_MetadataOptions) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.MetadataOptions"
}
