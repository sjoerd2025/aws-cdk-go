package awsagentregistry


// URL-based source configuration for a source-only descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceOnlyDescriptorSourceFromUrlProperty := &SourceOnlyDescriptorSourceFromUrlProperty{
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-sourceonlydescriptorsourcefromurl.html
//
type CfnRegistryRecord_SourceOnlyDescriptorSourceFromUrlProperty struct {
	// URL source for descriptor content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-sourceonlydescriptorsourcefromurl.html#cfn-agentregistry-registryrecord-sourceonlydescriptorsourcefromurl-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

