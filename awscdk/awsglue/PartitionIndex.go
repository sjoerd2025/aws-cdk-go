package awsglue


// Properties of a Partition Index.
//
// Example:
//   var myTable S3Table
//
//   myTable.AddPartitionIndex(&PartitionIndex{
//   	IndexName: jsii.String("my-index"),
//   	KeyNames: []*string{
//   		jsii.String("year"),
//   	},
//   })
//
type PartitionIndex struct {
	// The partition key names that comprise the partition index.
	//
	// The names must correspond to a name in the
	// table's partition keys.
	KeyNames *[]*string `field:"required" json:"keyNames" yaml:"keyNames"`
	// The name of the partition index.
	// Default: - a name will be generated for you.
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}

