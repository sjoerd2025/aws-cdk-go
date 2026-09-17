package awscdkgluealpha


// Properties for defining a Workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   workflowProps := &WorkflowProps{
//   	DefaultRunProperties: map[string]*string{
//   		"defaultRunPropertiesKey": jsii.String("defaultRunProperties"),
//   	},
//   	Description: jsii.String("description"),
//   	MaxConcurrentRuns: jsii.Number(123),
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
// Deprecated.
type WorkflowProps struct {
	// A map of properties to use when this workflow is executed.
	// Default: - no default run properties.
	//
	// Deprecated.
	DefaultRunProperties *map[string]*string `field:"optional" json:"defaultRunProperties" yaml:"defaultRunProperties"`
	// A description of the workflow.
	// Default: - no description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of concurrent runs allowed for the workflow.
	// Default: - no limit.
	//
	// Deprecated.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Name of the workflow.
	// Default: - a name will be generated.
	//
	// Deprecated.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

