package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecretsmanager"
)

// Base Connection Options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectionNetwork ConnectionNetwork
//   var secretRef ISecretRef
//   var securityGroup SecurityGroup
//
//   connectionOptions := &ConnectionOptions{
//   	ConnectionName: jsii.String("connectionName"),
//   	Description: jsii.String("description"),
//   	MatchCriteria: []*string{
//   		jsii.String("matchCriteria"),
//   	},
//   	Network: connectionNetwork,
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Secret: secretRef,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type ConnectionOptions struct {
	// The name of the connection.
	// Default: cloudformation generated name.
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The description of the connection.
	// Default: no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// This is useful for filtering the results of https://awscli.amazonaws.com/v2/documentation/api/latest/reference/glue/get-connections.html
	// Default: no match criteria.
	//
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// The VPC network placement for this connection, so it can reach resources inside a VPC. See more at https://docs.aws.amazon.com/glue/latest/dg/start-connecting.html.
	//
	// Build it with `ConnectionNetwork.subnet(subnet)` to pin a specific subnet,
	// or `ConnectionNetwork.vpc(vpc, vpcSubnets?)` to let the CDK select one.
	// Default: - no VPC network placement.
	//
	Network ConnectionNetwork `field:"optional" json:"network" yaml:"network"`
	// Key-Value pairs that define parameters for the connection.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-connect.html
	//
	// Default: empty properties.
	//
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// A reference to a Secrets Manager secret holding the credentials for this connection.
	//
	// The secret is referenced through the connection's `SECRET_ID` property, so
	// Glue reads the credentials at runtime and the secret value never appears in
	// the synthesized template. Prefer this over placing credentials directly in
	// `properties`. Accepts any `secretsmanager.ISecret`.
	// Default: - no secret; any credentials must be supplied via `properties`.
	//
	Secret interfacesawssecretsmanager.ISecretRef `field:"optional" json:"secret" yaml:"secret"`
	// The list of security groups needed to successfully make this connection e.g. to successfully connect to VPC.
	// Default: no security group.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

