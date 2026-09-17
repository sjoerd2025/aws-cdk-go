package awscdkgluealpha


// Job states emitted by Glue to CloudWatch Events.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types for more information.
//
// Deprecated.
type JobState string

const (
	// State indicating job run succeeded.
	// Deprecated.
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	// State indicating job run failed.
	// Deprecated.
	JobState_FAILED JobState = "FAILED"
	// State indicating job run timed out.
	// Deprecated.
	JobState_TIMEOUT JobState = "TIMEOUT"
	// State indicating job is starting.
	// Deprecated.
	JobState_STARTING JobState = "STARTING"
	// State indicating job is running.
	// Deprecated.
	JobState_RUNNING JobState = "RUNNING"
	// State indicating job is stopping.
	// Deprecated.
	JobState_STOPPING JobState = "STOPPING"
	// State indicating job stopped.
	// Deprecated.
	JobState_STOPPED JobState = "STOPPED"
)

