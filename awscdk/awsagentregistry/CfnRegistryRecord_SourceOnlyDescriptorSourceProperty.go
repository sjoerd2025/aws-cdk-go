package awsagentregistry


// Source configuration for a source-only descriptor.
//
// Unlike mcpServer/a2aAgentCard sources, source-only descriptors do not support credential providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceOnlyDescriptorSourceProperty := &SourceOnlyDescriptorSourceProperty{
//   	FromUrl: &SourceOnlyDescriptorSourceFromUrlProperty{
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-sourceonlydescriptorsource.html
//
type CfnRegistryRecord_SourceOnlyDescriptorSourceProperty struct {
	// URL-based source configuration for a source-only descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-sourceonlydescriptorsource.html#cfn-agentregistry-registryrecord-sourceonlydescriptorsource-fromurl
	//
	FromUrl interface{} `field:"optional" json:"fromUrl" yaml:"fromUrl"`
}

