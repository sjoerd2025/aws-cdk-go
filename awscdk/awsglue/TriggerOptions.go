package awsglue


// Properties for configuring a Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action Action
//
//   triggerOptions := &TriggerOptions{
//   	Actions: []Action{
//   		action,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
type TriggerOptions struct {
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
}

