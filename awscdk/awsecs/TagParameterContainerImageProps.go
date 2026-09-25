package awsecs


// Properties for `TagParameterContainerImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagParameterContainerImageProps := &TagParameterContainerImageProps{
//   	IsImageDigest: jsii.Boolean(false),
//   }
//
type TagParameterContainerImageProps struct {
	// Whether the CloudFormation Parameter holds an image digest (`sha256:...`) rather than a tag.
	//
	// When `true`, the separator between the repository URI and the parameter value is `@`
	// instead of `:`, producing `ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPO@sha256:...`.
	//
	// Use this when your pipeline passes a digest rather than a mutable tag.
	// Default: false.
	//
	IsImageDigest *bool `field:"optional" json:"isImageDigest" yaml:"isImageDigest"`
}

