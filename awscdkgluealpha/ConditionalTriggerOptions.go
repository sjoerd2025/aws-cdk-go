package awscdkgluealpha


// Properties for configuring a Condition (Predicate) based Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
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
//   		Logical: glue_alpha.PredicateLogical_AND,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StartOnCreation: jsii.Boolean(false),
//   }
//
// Deprecated.
type ConditionalTriggerOptions struct {
	// The actions initiated by this trigger.
	// Deprecated.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
	// A description for the trigger.
	// Default: - no description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the trigger.
	// Default: - no name is provided.
	//
	// Deprecated.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether to start the trigger on creation or not.
	// Default: - false.
	//
	// Deprecated.
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The predicate for the trigger.
	// Deprecated.
	Predicate *Predicate `field:"required" json:"predicate" yaml:"predicate"`
}

