package awsglue


// Construction properties for a `Catalog`.
//
// Example:
//   glue.NewCatalog(this, jsii.String("MyCatalog"), &CatalogProps{
//   	CatalogName: jsii.String("my-catalog"),
//   	Description: jsii.String("my catalog description"),
//   })
//
type CatalogProps struct {
	// Connection-password encryption configuration for the catalog.
	// Default: - connection-password encryption is not managed by CDK.
	//
	ConnectionPasswordEncryption *ConnectionPasswordEncryption `field:"optional" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// Encryption-at-rest configuration for the catalog.
	// Default: - encryption at rest is not managed by CDK (the catalog default applies).
	//
	EncryptionAtRest DataCatalogEncryptionAtRest `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// The name of the catalog.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// A description of the catalog.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

