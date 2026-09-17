package awsglue


// Specifies how to map columns when the table uses ORC data format.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"orc.schema.resolution"_
//
type OrcColumnMappingType string

const (
	// Map columns by name.
	OrcColumnMappingType_NAME OrcColumnMappingType = "NAME"
	// Map columns by position.
	OrcColumnMappingType_POSITION OrcColumnMappingType = "POSITION"
)

