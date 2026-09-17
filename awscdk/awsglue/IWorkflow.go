package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// The base interface for Glue Workflow.
// See: https://docs.aws.amazon.com/glue/latest/dg/workflows_overview.html
//
type IWorkflow interface {
	awscdk.IResource
	// Add a conditional (predicate-based) trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	AddConditionalTrigger(id *string, options *ConditionalTriggerOptions) interfacesawsglue.ITriggerRef
	// Add an EventBridge event-based trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	AddEventTrigger(id *string, options *EventTriggerOptions) interfacesawsglue.ITriggerRef
	// Add an on-demand trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) interfacesawsglue.ITriggerRef
	// Add a scheduled trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	AddScheduledTrigger(id *string, options *ScheduledTriggerOptions) interfacesawsglue.ITriggerRef
	// The ARN of the workflow.
	WorkflowArn() *string
	// The name of the workflow.
	WorkflowName() *string
}

// The jsii proxy for IWorkflow
type jsiiProxy_IWorkflow struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IWorkflow) AddConditionalTrigger(id *string, options *ConditionalTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddConditionalTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addConditionalTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddEventTrigger(id *string, options *EventTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddEventTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addEventTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddOnDemandTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addOnDemandTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddScheduledTrigger(id *string, options *ScheduledTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddScheduledTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addScheduledTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IWorkflow) WorkflowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflow) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}

