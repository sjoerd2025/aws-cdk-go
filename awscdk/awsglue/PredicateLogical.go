package awsglue


type PredicateLogical string

const (
	// All conditions must be true for the predicate to be true.
	PredicateLogical_AND PredicateLogical = "AND"
	// At least one condition must be true for the predicate to be true.
	PredicateLogical_ANY PredicateLogical = "ANY"
)

