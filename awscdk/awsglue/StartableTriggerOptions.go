package awsglue


// Base options for triggers that can be started when they are created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action Action
//
//   startableTriggerOptions := &StartableTriggerOptions{
//   	Actions: []Action{
//   		action,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
type StartableTriggerOptions struct {
	// The actions initiated by this trigger.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
	// A description for the trigger.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the trigger.
	// Default: - no name is provided.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether to start the trigger on creation or not.
	// Default: - false.
	//
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
}

