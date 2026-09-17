package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("data"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []Column{
//   		&Column{
//   			Name: jsii.String("date"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"date": glue.PartitionProjectionConfiguration_date(&DatePartitionProjectionConfigurationProps{
//   			"min": jsii.String("NOW-3YEARS"),
//   			"max": jsii.String("NOW"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/data-types.html
//
type Schema interface {
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

func NewSchema() Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.Schema",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSchema_Override(s Schema) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.Schema",
		nil, // no parameters
		s,
	)
}

// Creates an array of some other type.
func Schema_Array(itemType Type) Type {
	_init_.Initialize()

	if err := validateSchema_ArrayParameters(itemType); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"array",
		[]interface{}{itemType},
		&returns,
	)

	return returns
}

// Fixed length character data, with a specified length between 1 and 255.
func Schema_Char(length *float64) Type {
	_init_.Initialize()

	if err := validateSchema_CharParameters(length); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"char",
		[]interface{}{length},
		&returns,
	)

	return returns
}

// Creates a custom type from a raw Glue input string.
//
// Escape hatch for column types the other `Schema` factories don't model. The
// `inputString` is emitted verbatim and is not validated.
func Schema_Custom(inputString *string, isPrimitive *bool) Type {
	_init_.Initialize()

	if err := validateSchema_CustomParameters(inputString); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"custom",
		[]interface{}{inputString, isPrimitive},
		&returns,
	)

	return returns
}

// Creates a decimal type.
// See: https://docs.aws.amazon.com/athena/latest/ug/data-types.html
//
func Schema_Decimal(precision *float64, scale *float64) Type {
	_init_.Initialize()

	if err := validateSchema_DecimalParameters(precision); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"decimal",
		[]interface{}{precision, scale},
		&returns,
	)

	return returns
}

// Creates a map of some primitive key type to some value type.
func Schema_Map(keyType Type, valueType Type) Type {
	_init_.Initialize()

	if err := validateSchema_MapParameters(keyType, valueType); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"map",
		[]interface{}{keyType, valueType},
		&returns,
	)

	return returns
}

// Creates a nested structure containing individually named and typed columns.
func Schema_Struct(columns *[]*Column) Type {
	_init_.Initialize()

	if err := validateSchema_StructParameters(columns); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"struct",
		[]interface{}{columns},
		&returns,
	)

	return returns
}

// Variable length character data, with a specified length between 1 and 65535.
func Schema_Varchar(length *float64) Type {
	_init_.Initialize()

	if err := validateSchema_VarcharParameters(length); err != nil {
		panic(err)
	}
	var returns Type

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.Schema",
		"varchar",
		[]interface{}{length},
		&returns,
	)

	return returns
}

func Schema_BIG_INT() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"BIG_INT",
		&returns,
	)
	return returns
}

func Schema_BINARY() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"BINARY",
		&returns,
	)
	return returns
}

func Schema_BOOLEAN() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"BOOLEAN",
		&returns,
	)
	return returns
}

func Schema_DATE() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"DATE",
		&returns,
	)
	return returns
}

func Schema_DOUBLE() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"DOUBLE",
		&returns,
	)
	return returns
}

func Schema_FLOAT() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"FLOAT",
		&returns,
	)
	return returns
}

func Schema_INTEGER() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"INTEGER",
		&returns,
	)
	return returns
}

func Schema_SMALL_INT() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"SMALL_INT",
		&returns,
	)
	return returns
}

func Schema_STRING() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"STRING",
		&returns,
	)
	return returns
}

func Schema_TIMESTAMP() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"TIMESTAMP",
		&returns,
	)
	return returns
}

func Schema_TINY_INT() Type {
	_init_.Initialize()
	var returns Type
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.Schema",
		"TINY_INT",
		&returns,
	)
	return returns
}

