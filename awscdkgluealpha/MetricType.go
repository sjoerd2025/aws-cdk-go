package awscdkgluealpha


// The Glue CloudWatch metric type.
// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
//
// Deprecated.
type MetricType string

const (
	// A value at a point in time.
	// Deprecated.
	MetricType_GAUGE MetricType = "GAUGE"
	// An aggregate number.
	// Deprecated.
	MetricType_COUNT MetricType = "COUNT"
)

