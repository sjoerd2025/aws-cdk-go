package awsglue


// Job states emitted by Glue to CloudWatch Events.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types for more information.
//
type JobState string

const (
	// State indicating job run succeeded.
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	// State indicating job run failed.
	JobState_FAILED JobState = "FAILED"
	// State indicating job run timed out.
	JobState_TIMEOUT JobState = "TIMEOUT"
	// State indicating job is starting.
	JobState_STARTING JobState = "STARTING"
	// State indicating job is running.
	JobState_RUNNING JobState = "RUNNING"
	// State indicating job is stopping.
	JobState_STOPPING JobState = "STOPPING"
	// State indicating job stopped.
	JobState_STOPPED JobState = "STOPPED"
)

