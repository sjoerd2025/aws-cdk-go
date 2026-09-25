package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Optional parameters for creating code using bucket.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   var key Key
//
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   options := map[string]Key{
//   	"sourceKMSKey": key,
//   }
//   fnBucket := lambda.NewFunction(this, jsii.String("myFunction2"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromBucketV2(bucket, jsii.String("python-lambda-handler.zip"), options),
//   })
//
type BucketOptions struct {
	// Optional S3 object version.
	//
	// Required when `s3ObjectStorageMode` is set to `S3ObjectStorageMode.REFERENCE`.
	// Default: - no object version.
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
	// How Lambda manages the storage of your code package.
	// Default: - Lambda copies the deployment package from your S3 bucket into Lambda-managed storage.
	//
	S3ObjectStorageMode S3ObjectStorageMode `field:"optional" json:"s3ObjectStorageMode" yaml:"s3ObjectStorageMode"`
	// The KMS key that Lambda uses to encrypt the deployment package in Lambda-managed storage.
	//
	// This is not the key used to encrypt the source object in Amazon S3.
	// Default: - Lambda uses an AWS owned key.
	//
	SourceKMSKey interfacesawskms.IKeyRef `field:"optional" json:"sourceKMSKey" yaml:"sourceKMSKey"`
}

