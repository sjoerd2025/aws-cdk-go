package awscdkgluealpha


// Represents the state of a crawler for a condition in the Glue Trigger API.
// Deprecated.
type CrawlerState string

const (
	// The crawler is currently running.
	// Deprecated.
	CrawlerState_RUNNING CrawlerState = "RUNNING"
	// The crawler is in the process of being cancelled.
	// Deprecated.
	CrawlerState_CANCELLING CrawlerState = "CANCELLING"
	// The crawler has been cancelled.
	// Deprecated.
	CrawlerState_CANCELLED CrawlerState = "CANCELLED"
	// The crawler has completed its operation successfully.
	// Deprecated.
	CrawlerState_SUCCEEDED CrawlerState = "SUCCEEDED"
	// The crawler has failed to complete its operation.
	// Deprecated.
	CrawlerState_FAILED CrawlerState = "FAILED"
	// The crawler encountered an error during its operation.
	// Deprecated.
	CrawlerState_ERROR CrawlerState = "ERROR"
)

