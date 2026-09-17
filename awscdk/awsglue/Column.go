package awsglue


// A column of a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var type Type
//
//   column := &Column{
//   	Name: jsii.String("name"),
//   	Type: type,
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
type Column struct {
	// Name of the column.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the column.
	Type Type `field:"required" json:"type" yaml:"type"`
	// Coment describing the column.
	// Default: none.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

