package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
)

type IDataQualityRuleset interface {
	awscdk.IResource
	// The ARN of the ruleset.
	RulesetArn() *string
	// The name of the ruleset.
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

