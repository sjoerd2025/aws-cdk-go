package awscdkgluealpha


// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARCHAR, CHAR, or string data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"surplus_char_handling"_
//
// Deprecated.
type SurplusCharHandlingAction string

const (
	// Replaces data that exceeds the column width with null.
	// Deprecated.
	SurplusCharHandlingAction_SET_TO_NULL SurplusCharHandlingAction = "SET_TO_NULL"
	// Doesn't perform surplus character handling.
	// Deprecated.
	SurplusCharHandlingAction_DISABLED SurplusCharHandlingAction = "DISABLED"
	// Cancels queries that return data exceeding the column width.
	// Deprecated.
	SurplusCharHandlingAction_FAIL SurplusCharHandlingAction = "FAIL"
	// Replaces each value in the row with null.
	// Deprecated.
	SurplusCharHandlingAction_DROP_ROW SurplusCharHandlingAction = "DROP_ROW"
	// Removes the characters that exceed the maximum number of characters defined for the column.
	// Deprecated.
	SurplusCharHandlingAction_TRUNCATE SurplusCharHandlingAction = "TRUNCATE"
)

