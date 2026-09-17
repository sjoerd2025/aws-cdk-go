package awscdkgluealpha


// A column of a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
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
// Deprecated.
type Column struct {
	// Name of the column.
	// Deprecated.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the column.
	// Deprecated.
	Type Type `field:"required" json:"type" yaml:"type"`
	// Coment describing the column.
	// Default: none.
	//
	// Deprecated.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

