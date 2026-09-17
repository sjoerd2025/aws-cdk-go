package awsglue


// Properties for configuring a Condition (Predicate) based Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action Action
//   var condition Condition
//
//   conditionalTriggerOptions := &ConditionalTriggerOptions{
//   	Actions: []Action{
//   		action,
//   	},
//   	Predicate: &Predicate{
//   		Conditions: []Condition{
//   			condition,
//   		},
//   		Logical: awscdk.Aws_glue.PredicateLogical_AND,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
type ConditionalTriggerOptions struct {
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
	// The predicate for the trigger.
	Predicate *Predicate `field:"required" json:"predicate" yaml:"predicate"`
}

