package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The Spark UI logging location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   sparkUILoggingLocation := &SparkUILoggingLocation{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
type SparkUILoggingLocation struct {
	// The bucket where the Glue job stores the logs.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Default: '/' - the logs will be written at the root of the bucket.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

