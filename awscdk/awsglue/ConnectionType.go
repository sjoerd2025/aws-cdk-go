package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of the glue connection.
//
// If you need to use a connection type that doesn't exist as a static member, you
// can instantiate a `ConnectionType` object, e.g: `new ConnectionType('NEW_TYPE')`.
//
// Example:
//   var securityGroup SecurityGroup
//   var vpc Vpc
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
//   	Type: glue.ConnectionType_NETWORK(),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	// vpcSubnets is optional - defaults to private subnets
//   	Network: glue.ConnectionNetwork_Vpc(vpc, &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype
//
type ConnectionType interface {
	// The name of this ConnectionType, as expected by Connection resource.
	Name() *string
	// The connection type name as expected by Connection resource.
	ToString() *string
}

// The jsii proxy struct for ConnectionType
type jsiiProxy_ConnectionType struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewConnectionType(name *string) ConnectionType {
	_init_.Initialize()

	if err := validateNewConnectionTypeParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectionType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.ConnectionType",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewConnectionType_Override(c ConnectionType, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.ConnectionType",
		[]interface{}{name},
		c,
	)
}

func ConnectionType_AZURECOSMOS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"AZURECOSMOS",
		&returns,
	)
	return returns
}

func ConnectionType_AZURESQL() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"AZURESQL",
		&returns,
	)
	return returns
}

func ConnectionType_BIGQUERY() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"BIGQUERY",
		&returns,
	)
	return returns
}

func ConnectionType_CUSTOM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"CUSTOM",
		&returns,
	)
	return returns
}

func ConnectionType_DYNAMODB() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"DYNAMODB",
		&returns,
	)
	return returns
}

func ConnectionType_FACEBOOKADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"FACEBOOKADS",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLEADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"GOOGLEADS",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLEANALYTICS4() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"GOOGLEANALYTICS4",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLESHEETS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"GOOGLESHEETS",
		&returns,
	)
	return returns
}

func ConnectionType_HUBSPOT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"HUBSPOT",
		&returns,
	)
	return returns
}

func ConnectionType_INSTAGRAMADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"INSTAGRAMADS",
		&returns,
	)
	return returns
}

func ConnectionType_INTERCOM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"INTERCOM",
		&returns,
	)
	return returns
}

func ConnectionType_JDBC() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"JDBC",
		&returns,
	)
	return returns
}

func ConnectionType_JIRACLOUD() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"JIRACLOUD",
		&returns,
	)
	return returns
}

func ConnectionType_KAFKA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"KAFKA",
		&returns,
	)
	return returns
}

func ConnectionType_MARKETO() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"MARKETO",
		&returns,
	)
	return returns
}

func ConnectionType_MARKETPLACE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"MARKETPLACE",
		&returns,
	)
	return returns
}

func ConnectionType_MONGODB() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"MONGODB",
		&returns,
	)
	return returns
}

func ConnectionType_MYSQL() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"MYSQL",
		&returns,
	)
	return returns
}

func ConnectionType_NETSUITEERP() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"NETSUITEERP",
		&returns,
	)
	return returns
}

func ConnectionType_NETWORK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"NETWORK",
		&returns,
	)
	return returns
}

func ConnectionType_OPENSEARCH() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"OPENSEARCH",
		&returns,
	)
	return returns
}

func ConnectionType_ORACLE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"ORACLE",
		&returns,
	)
	return returns
}

func ConnectionType_POSTGRESQL() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"POSTGRESQL",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SALESFORCE",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCEMARKETINGCLOUD() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SALESFORCEMARKETINGCLOUD",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCEPARDOT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SALESFORCEPARDOT",
		&returns,
	)
	return returns
}

func ConnectionType_SAPHANA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SAPHANA",
		&returns,
	)
	return returns
}

func ConnectionType_SAPODATA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SAPODATA",
		&returns,
	)
	return returns
}

func ConnectionType_SERVICENOW() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SERVICENOW",
		&returns,
	)
	return returns
}

func ConnectionType_SLACK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SLACK",
		&returns,
	)
	return returns
}

func ConnectionType_SNAPCHATADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SNAPCHATADS",
		&returns,
	)
	return returns
}

func ConnectionType_SQLSERVER() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"SQLSERVER",
		&returns,
	)
	return returns
}

func ConnectionType_STRIPE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"STRIPE",
		&returns,
	)
	return returns
}

func ConnectionType_TERADATA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"TERADATA",
		&returns,
	)
	return returns
}

func ConnectionType_VERTICA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"VERTICA",
		&returns,
	)
	return returns
}

func ConnectionType_VIEW_VALIDATION_ATHENA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"VIEW_VALIDATION_ATHENA",
		&returns,
	)
	return returns
}

func ConnectionType_VIEW_VALIDATION_REDSHIFT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"VIEW_VALIDATION_REDSHIFT",
		&returns,
	)
	return returns
}

func ConnectionType_ZENDESK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"ZENDESK",
		&returns,
	)
	return returns
}

func ConnectionType_ZOHOCRM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.ConnectionType",
		"ZOHOCRM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

