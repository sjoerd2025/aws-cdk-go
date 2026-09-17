package awsglue


// Represents a trigger predicate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var condition Condition
//
//   predicate := &Predicate{
//   	Conditions: []Condition{
//   		condition,
//   	},
//   	Logical: awscdk.Aws_glue.PredicateLogical_AND,
//   }
//
type Predicate struct {
	// A list of the conditions that determine when the trigger will fire.
	// Default: - no conditions are provided.
	//
	Conditions *[]Condition `field:"optional" json:"conditions" yaml:"conditions"`
	// The logical operator to be applied to the conditions.
	// Default: - PredicateLogical.AND if multiple conditions are provided, no logical operator if only one condition
	//
	Logical PredicateLogical `field:"optional" json:"logical" yaml:"logical"`
}

