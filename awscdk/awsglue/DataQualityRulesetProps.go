package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for `DataQualityRuleset`.
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
type DataQualityRulesetProps struct {
	// The DQDL document defining the ruleset's data quality rules.
	//
	// Build it with `Dqdl.fromString(...)`.
	Dqdl Dqdl `field:"required" json:"dqdl" yaml:"dqdl"`
	// The name of the ruleset.
	RulesetName *string `field:"required" json:"rulesetName" yaml:"rulesetName"`
	// The target table of the ruleset.
	TargetTable DataQualityTargetTable `field:"required" json:"targetTable" yaml:"targetTable"`
	// The description of the ruleset.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Policy to apply when the ruleset is removed from the stack.
	// Default: - resource will be destroyed.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Key-Value pairs that define tags for the ruleset.
	// Default: empty tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

