package awscdkgluealpha


// Properties for importing a Workflow using its attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   workflowAttributes := &WorkflowAttributes{
//   	WorkflowName: jsii.String("workflowName"),
//
//   	// the properties below are optional
//   	WorkflowArn: jsii.String("workflowArn"),
//   }
//
// Deprecated.
type WorkflowAttributes struct {
	// The name of the workflow to import.
	// Deprecated.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The ARN of the workflow to import.
	// Default: - derived from the workflow name.
	//
	// Deprecated.
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
}

