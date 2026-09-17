package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableAttributes := &TableAttributes{
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
type TableAttributes struct {
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

