package awsglue


// Represents the logical operator for evaluating a single condition in the Glue Trigger API.
type ConditionLogicalOperator string

const (
	// The condition is true if the values are equal.
	ConditionLogicalOperator_EQUALS ConditionLogicalOperator = "EQUALS"
)

