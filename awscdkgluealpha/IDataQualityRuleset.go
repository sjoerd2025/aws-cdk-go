package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// Deprecated.
type IDataQualityRuleset interface {
	awscdk.IResource
	// The ARN of the ruleset.
	// Deprecated.
	RulesetArn() *string
	// The name of the ruleset.
	// Deprecated.
	RulesetName() *string
}

// The jsii proxy for IDataQualityRuleset
type jsiiProxy_IDataQualityRuleset struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDataQualityRuleset) RulesetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataQualityRuleset) RulesetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesetName",
		&returns,
	)
	return returns
}

