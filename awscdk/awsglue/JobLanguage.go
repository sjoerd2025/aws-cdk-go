package awsglue


// Runtime language of the Glue job.
type JobLanguage string

const (
	// Scala.
	JobLanguage_SCALA JobLanguage = "SCALA"
	// Python.
	JobLanguage_PYTHON JobLanguage = "PYTHON"
)

