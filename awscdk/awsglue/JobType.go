package awsglue


// The job type.
type JobType string

const (
	// Command for running a Glue Spark job.
	JobType_ETL JobType = "ETL"
	// Command for running a Glue Spark streaming job.
	JobType_STREAMING JobType = "STREAMING"
	// Command for running a Glue python shell job.
	JobType_PYTHON_SHELL JobType = "PYTHON_SHELL"
)

