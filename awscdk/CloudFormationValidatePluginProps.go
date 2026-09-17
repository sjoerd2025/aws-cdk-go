package awscdk


// Properties for configuring the CloudFormationValidatePlugin.
//
// Example:
//   // Rules text, read from disk perhaps
//   var myRules string
//   app := awscdk.NewApp()
//
//   awscdk.Validations_Of(app).AddPlugins(awscdk.NewCloudFormationValidatePlugin(&CloudFormationValidatePluginProps{
//   	GuardRules: []ValidationRuleSource{
//   		&ValidationRuleSource{
//   			Name: jsii.String("My rules"),
//   			Content: myRules,
//   		},
//   	},
//   }))
//
type CloudFormationValidatePluginProps struct {
	// Custom Guard rules to evaluate in addition to built-in rules.
	// Default: - no guard rules.
	//
	GuardRules *[]*ValidationRuleSource `field:"optional" json:"guardRules" yaml:"guardRules"`
	// Whether to evaluate the default Rego rules that ship with the CDK.
	//
	// Registering a `CloudFormationValidatePlugin` explicitly replaces the
	// auto-registered default instance, so without this flag adding custom
	// rules would silently drop the CDK default rules. Individual default
	// rules can be suppressed by ID via `Validations.of(scope).acknowledge()`.
	// Default: true.
	//
	IncludeDefaultRules *bool `field:"optional" json:"includeDefaultRules" yaml:"includeDefaultRules"`
	// Custom Rego rules to evaluate in addition to built-in rules.
	// Default: - no custom rules.
	//
	RegoRules *[]*ValidationRuleSource `field:"optional" json:"regoRules" yaml:"regoRules"`
}

