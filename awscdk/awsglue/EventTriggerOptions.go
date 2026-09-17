package awsglue


// Properties for configuring an Event Bridge based Glue Trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action Action
//
//   eventTriggerOptions := &EventTriggerOptions{
//   	Actions: []Action{
//   		action,
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EventBatchingCondition: &EventBatchingCondition{
//   		BatchSize: jsii.Number(123),
//
//   		// the properties below are optional
//   		BatchWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	},
//   	Name: jsii.String("name"),
//   }
//
type EventTriggerOptions struct {
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
	// Batch condition for the trigger.
	// Default: - no batch condition.
	//
	EventBatchingCondition *EventBatchingCondition `field:"optional" json:"eventBatchingCondition" yaml:"eventBatchingCondition"`
}

