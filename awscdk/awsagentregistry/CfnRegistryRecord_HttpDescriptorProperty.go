package awsagentregistry


// The HTTP descriptor, populated for records detected from an HTTP protocol source.
//
// This descriptor is source-only: its content is synchronized from the configured source URL rather than supplied inline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpDescriptorProperty := &HttpDescriptorProperty{
//   	Source: &SourceOnlyDescriptorSourceProperty{
//   		FromUrl: &SourceOnlyDescriptorSourceFromUrlProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-httpdescriptor.html
//
type CfnRegistryRecord_HttpDescriptorProperty struct {
	// Source configuration for a source-only descriptor.
	//
	// Unlike mcpServer/a2aAgentCard sources, source-only descriptors do not support credential providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-httpdescriptor.html#cfn-agentregistry-registryrecord-httpdescriptor-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

