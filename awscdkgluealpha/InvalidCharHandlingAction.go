package awscdkgluealpha


// Specifies the action to perform when query results contain invalid UTF-8 character values.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"invalid_char_handling"_
//
// Deprecated.
type InvalidCharHandlingAction string

const (
	// Doesn't perform invalid character handling.
	// Deprecated.
	InvalidCharHandlingAction_DISABLED InvalidCharHandlingAction = "DISABLED"
	// Cancels queries that return data containing invalid UTF-8 values.
	// Deprecated.
	InvalidCharHandlingAction_FAIL InvalidCharHandlingAction = "FAIL"
	// Replaces invalid UTF-8 values with null.
	// Deprecated.
	InvalidCharHandlingAction_SET_TO_NULL InvalidCharHandlingAction = "SET_TO_NULL"
	// Replaces each value in the row with null.
	// Deprecated.
	InvalidCharHandlingAction_DROP_ROW InvalidCharHandlingAction = "DROP_ROW"
	// Replaces the invalid character with the replacement character you specify using `REPLACEMENT_CHAR`.
	// Deprecated.
	InvalidCharHandlingAction_REPLACE InvalidCharHandlingAction = "REPLACE"
)

