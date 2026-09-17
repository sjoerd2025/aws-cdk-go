package awsglue


// Partition projection type.
//
// Determines how Athena projects partition values.
// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html
//
type PartitionProjectionType string

const (
	// Project partition values as integers within a range.
	PartitionProjectionType_INTEGER PartitionProjectionType = "INTEGER"
	// Project partition values as dates within a range.
	PartitionProjectionType_DATE PartitionProjectionType = "DATE"
	// Project partition values from an explicit list of values.
	PartitionProjectionType_ENUM PartitionProjectionType = "ENUM"
	// Project partition values that are injected at query time.
	PartitionProjectionType_INJECTED PartitionProjectionType = "INJECTED"
)

