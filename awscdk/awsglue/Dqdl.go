package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Data Quality Definition Language (DQDL) document for a `DataQualityRuleset`.
//
// DQDL is an authored string that Glue parses and validates at deploy time. Build
// one from a raw DQDL string with {@link Dqdl.fromString}.
//
// Example:
//   var database IDatabase
//
//   glue.NewDataQualityRuleset(this, jsii.String("MyRuleset"), &DataQualityRulesetProps{
//   	RulesetName: jsii.String("my_ruleset"),
//   	Dqdl: glue.Dqdl_FromString(jsii.String("Rules = [ RowCount > 100, IsComplete \"order_id\" ]")),
//   	TargetTable: glue.DataQualityTargetTable_FromTableName(database, jsii.String("my_table")),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/dg/dqdl.html
//
type Dqdl interface {
}

// The jsii proxy struct for Dqdl
type jsiiProxy_Dqdl struct {
	_ byte // padding
}

// Create a `Dqdl` from a raw DQDL string.
func Dqdl_FromString(dqdl *string) Dqdl {
	_init_.Initialize()

	if err := validateDqdl_FromStringParameters(dqdl); err != nil {
		panic(err)
	}
	var returns Dqdl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Dqdl",
		"fromString",
		[]interface{}{dqdl},
		&returns,
	)

	return returns
}

