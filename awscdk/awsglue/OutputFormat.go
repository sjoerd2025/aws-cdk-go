package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Absolute class name of the Hadoop `OutputFormat` to use when writing table files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFormat := awscdk.Aws_glue.OutputFormat_AVRO()
//
type OutputFormat interface {
	ClassName() *string
}

// The jsii proxy struct for OutputFormat
type jsiiProxy_OutputFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_OutputFormat) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


func NewOutputFormat(className *string) OutputFormat {
	_init_.Initialize()

	if err := validateNewOutputFormatParameters(className); err != nil {
		panic(err)
	}
	j := jsiiProxy_OutputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.OutputFormat",
		[]interface{}{className},
		&j,
	)

	return &j
}

func NewOutputFormat_Override(o OutputFormat, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.OutputFormat",
		[]interface{}{className},
		o,
	)
}

func OutputFormat_AVRO() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.OutputFormat",
		"AVRO",
		&returns,
	)
	return returns
}

func OutputFormat_HIVE_IGNORE_KEY_TEXT() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.OutputFormat",
		"HIVE_IGNORE_KEY_TEXT",
		&returns,
	)
	return returns
}

func OutputFormat_ORC() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.OutputFormat",
		"ORC",
		&returns,
	)
	return returns
}

func OutputFormat_PARQUET() OutputFormat {
	_init_.Initialize()
	var returns OutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.OutputFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

