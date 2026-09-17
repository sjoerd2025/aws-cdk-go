package awsglue


// Options shared by all trigger conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionOptions := &ConditionOptions{
//   	LogicalOperator: awscdk.Aws_glue.ConditionLogicalOperator_EQUALS,
//   }
//
type ConditionOptions struct {
	// The logical operator for the condition.
	// Default: ConditionLogicalOperator.EQUALS
	//
	LogicalOperator ConditionLogicalOperator `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

