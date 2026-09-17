package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents event trigger batch condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBatchingCondition := &EventBatchingCondition{
//   	BatchSize: jsii.Number(123),
//
//   	// the properties below are optional
//   	BatchWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type EventBatchingCondition struct {
	// Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// Window of time in seconds after which EventBridge event trigger fires.
	// Default: - 900 seconds.
	//
	BatchWindow awscdk.Duration `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}

