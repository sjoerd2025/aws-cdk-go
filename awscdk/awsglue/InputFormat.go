package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Absolute class name of the Hadoop `InputFormat` to use when reading table files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputFormat := awscdk.Aws_glue.InputFormat_AVRO()
//
type InputFormat interface {
	ClassName() *string
}

// The jsii proxy struct for InputFormat
type jsiiProxy_InputFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_InputFormat) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


func NewInputFormat(className *string) InputFormat {
	_init_.Initialize()

	if err := validateNewInputFormatParameters(className); err != nil {
		panic(err)
	}
	j := jsiiProxy_InputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.InputFormat",
		[]interface{}{className},
		&j,
	)

	return &j
}

func NewInputFormat_Override(i InputFormat, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.InputFormat",
		[]interface{}{className},
		i,
	)
}

func InputFormat_AVRO() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.InputFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func InputFormat_CLOUDTRAIL() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.InputFormat",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func InputFormat_ORC() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.InputFormat",
		"ORC",
		&returns,
	)
	return returns
}

func InputFormat_PARQUET() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.InputFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

func InputFormat_TEXT() InputFormat {
	_init_.Initialize()
	var returns InputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.InputFormat",
		"TEXT",
		&returns,
	)
	return returns
}

