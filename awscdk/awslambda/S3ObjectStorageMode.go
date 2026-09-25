package awslambda


// How Lambda manages the storage of your code package.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//   var objectVersion string
//
//
//   lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	FunctionName: jsii.String("my-function"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromBucketV2(bucket, jsii.String("my-function.zip"), &BucketOptions{
//   		ObjectVersion: jsii.String(*ObjectVersion),
//   		S3ObjectStorageMode: lambda.S3ObjectStorageMode_REFERENCE,
//   	}),
//   })
//
//   lambda.NewLayerVersion(this, jsii.String("MyLayer"), &LayerVersionProps{
//   	LayerVersionName: jsii.String("my-layer"),
//   	Code: lambda.Code_*FromBucketV2(bucket, jsii.String("my-layer.zip"), &BucketOptions{
//   		ObjectVersion: jsii.String(*ObjectVersion),
//   		S3ObjectStorageMode: lambda.S3ObjectStorageMode_REFERENCE,
//   	}),
//   })
//
type S3ObjectStorageMode string

const (
	// Lambda copies the deployment package from your S3 bucket into Lambda-managed storage.
	S3ObjectStorageMode_COPY S3ObjectStorageMode = "COPY"
	// Lambda references your code directly from your S3 bucket.
	S3ObjectStorageMode_REFERENCE S3ObjectStorageMode = "REFERENCE"
)

