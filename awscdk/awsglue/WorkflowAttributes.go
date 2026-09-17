package awsglue


// Properties for importing a Workflow using its attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowAttributes := &WorkflowAttributes{
//   	WorkflowName: jsii.String("workflowName"),
//
//   	// the properties below are optional
//   	WorkflowArn: jsii.String("workflowArn"),
//   }
//
type WorkflowAttributes struct {
	// The name of the workflow to import.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The ARN of the workflow to import.
	// Default: - derived from the workflow name.
	//
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
}

