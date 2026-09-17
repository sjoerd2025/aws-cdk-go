package awsglue


// Represents the state of a crawler for a condition in the Glue Trigger API.
type CrawlerState string

const (
	// The crawler is currently running.
	CrawlerState_RUNNING CrawlerState = "RUNNING"
	// The crawler is in the process of being cancelled.
	CrawlerState_CANCELLING CrawlerState = "CANCELLING"
	// The crawler has been cancelled.
	CrawlerState_CANCELLED CrawlerState = "CANCELLED"
	// The crawler has completed its operation successfully.
	CrawlerState_SUCCEEDED CrawlerState = "SUCCEEDED"
	// The crawler has failed to complete its operation.
	CrawlerState_FAILED CrawlerState = "FAILED"
	// The crawler encountered an error during its operation.
	CrawlerState_ERROR CrawlerState = "ERROR"
)

