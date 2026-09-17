package awsglue


// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARBYTE data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"surplus_bytes_handling"_
//
type SurplusBytesHandlingAction string

const (
	// Replaces data that exceeds the column width with null.
	SurplusBytesHandlingAction_SET_TO_NULL SurplusBytesHandlingAction = "SET_TO_NULL"
	// Doesn't perform surplus byte handling.
	SurplusBytesHandlingAction_DISABLED SurplusBytesHandlingAction = "DISABLED"
	// Cancels queries that return data exceeding the column width.
	SurplusBytesHandlingAction_FAIL SurplusBytesHandlingAction = "FAIL"
	// Drop all rows that contain data exceeding column width.
	SurplusBytesHandlingAction_DROP_ROW SurplusBytesHandlingAction = "DROP_ROW"
	// Removes the characters that exceed the maximum number of characters defined for the column.
	SurplusBytesHandlingAction_TRUNCATE SurplusBytesHandlingAction = "TRUNCATE"
)

