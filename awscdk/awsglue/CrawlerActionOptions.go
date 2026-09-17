package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for the execution of a crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityConfiguration SecurityConfiguration
//
//   crawlerActionOptions := &CrawlerActionOptions{
//   	Arguments: map[string]*string{
//   		"argumentsKey": jsii.String("arguments"),
//   	},
//   	SecurityConfiguration: securityConfiguration,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type CrawlerActionOptions struct {
	// The arguments used when this trigger fires.
	// Default: - no arguments are passed to the job.
	//
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// The `SecurityConfiguration` to be used with this action.
	// Default: - no security configuration is used.
	//
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The run timeout.
	//
	// This is the maximum time that a run can consume resources before it is terminated and enters TIMEOUT status.
	// Default: - the default timeout value set in the job definition.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

