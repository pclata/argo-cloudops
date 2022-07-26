// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package testhelpers

import (
	"context"
	"github.com/cello-proj/cello/service/internal/workflow"
	"net/http"
	"sync"
)

// Ensure, that WorkflowMock does implement workflow.Workflow.
// If this is not the case, regenerate this file with moq.
var _ workflow.Workflow = &WorkflowMock{}

// WorkflowMock is a mock implementation of workflow.Workflow.
//
// 	func TestSomethingThatUsesWorkflow(t *testing.T) {
//
// 		// make and configure a mocked workflow.Workflow
// 		mockedWorkflow := &WorkflowMock{
// 			ListFunc: func(ctx context.Context) ([]workflow.Status, error) {
// 				panic("mock out the List method")
// 			},
// 			LogStreamFunc: func(ctx context.Context, workflowName string, data http.ResponseWriter) error {
// 				panic("mock out the LogStream method")
// 			},
// 			LogsFunc: func(ctx context.Context, workflowName string) (*workflow.Logs, error) {
// 				panic("mock out the Logs method")
// 			},
// 			StatusFunc: func(ctx context.Context, workflowName string) (*workflow.Status, error) {
// 				panic("mock out the Status method")
// 			},
// 			SubmitFunc: func(ctx context.Context, from string, parameters map[string]string, labels map[string]string) (string, error) {
// 				panic("mock out the Submit method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires workflow.Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context) ([]workflow.Status, error)

	// LogStreamFunc mocks the LogStream method.
	LogStreamFunc func(ctx context.Context, workflowName string, data http.ResponseWriter) error

	// LogsFunc mocks the Logs method.
	LogsFunc func(ctx context.Context, workflowName string) (*workflow.Logs, error)

	// StatusFunc mocks the Status method.
	StatusFunc func(ctx context.Context, workflowName string) (*workflow.Status, error)

	// SubmitFunc mocks the Submit method.
	SubmitFunc func(ctx context.Context, from string, parameters map[string]string, labels map[string]string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// LogStream holds details about calls to the LogStream method.
		LogStream []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WorkflowName is the workflowName argument value.
			WorkflowName string
			// Data is the data argument value.
			Data http.ResponseWriter
		}
		// Logs holds details about calls to the Logs method.
		Logs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WorkflowName is the workflowName argument value.
			WorkflowName string
		}
		// Status holds details about calls to the Status method.
		Status []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WorkflowName is the workflowName argument value.
			WorkflowName string
		}
		// Submit holds details about calls to the Submit method.
		Submit []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// From is the from argument value.
			From string
			// Parameters is the parameters argument value.
			Parameters map[string]string
			// Labels is the labels argument value.
			Labels map[string]string
		}
	}
	lockList      sync.RWMutex
	lockLogStream sync.RWMutex
	lockLogs      sync.RWMutex
	lockStatus    sync.RWMutex
	lockSubmit    sync.RWMutex
}

// List calls ListFunc.
func (mock *WorkflowMock) List(ctx context.Context) ([]workflow.Status, error) {
	if mock.ListFunc == nil {
		panic("WorkflowMock.ListFunc: method is nil but Workflow.List was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedWorkflow.ListCalls())
func (mock *WorkflowMock) ListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// LogStream calls LogStreamFunc.
func (mock *WorkflowMock) LogStream(ctx context.Context, workflowName string, data http.ResponseWriter) error {
	if mock.LogStreamFunc == nil {
		panic("WorkflowMock.LogStreamFunc: method is nil but Workflow.LogStream was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		WorkflowName string
		Data         http.ResponseWriter
	}{
		Ctx:          ctx,
		WorkflowName: workflowName,
		Data:         data,
	}
	mock.lockLogStream.Lock()
	mock.calls.LogStream = append(mock.calls.LogStream, callInfo)
	mock.lockLogStream.Unlock()
	return mock.LogStreamFunc(ctx, workflowName, data)
}

// LogStreamCalls gets all the calls that were made to LogStream.
// Check the length with:
//     len(mockedWorkflow.LogStreamCalls())
func (mock *WorkflowMock) LogStreamCalls() []struct {
	Ctx          context.Context
	WorkflowName string
	Data         http.ResponseWriter
} {
	var calls []struct {
		Ctx          context.Context
		WorkflowName string
		Data         http.ResponseWriter
	}
	mock.lockLogStream.RLock()
	calls = mock.calls.LogStream
	mock.lockLogStream.RUnlock()
	return calls
}

// Logs calls LogsFunc.
func (mock *WorkflowMock) Logs(ctx context.Context, workflowName string) (*workflow.Logs, error) {
	if mock.LogsFunc == nil {
		panic("WorkflowMock.LogsFunc: method is nil but Workflow.Logs was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		WorkflowName string
	}{
		Ctx:          ctx,
		WorkflowName: workflowName,
	}
	mock.lockLogs.Lock()
	mock.calls.Logs = append(mock.calls.Logs, callInfo)
	mock.lockLogs.Unlock()
	return mock.LogsFunc(ctx, workflowName)
}

// LogsCalls gets all the calls that were made to Logs.
// Check the length with:
//     len(mockedWorkflow.LogsCalls())
func (mock *WorkflowMock) LogsCalls() []struct {
	Ctx          context.Context
	WorkflowName string
} {
	var calls []struct {
		Ctx          context.Context
		WorkflowName string
	}
	mock.lockLogs.RLock()
	calls = mock.calls.Logs
	mock.lockLogs.RUnlock()
	return calls
}

// Status calls StatusFunc.
func (mock *WorkflowMock) Status(ctx context.Context, workflowName string) (*workflow.Status, error) {
	if mock.StatusFunc == nil {
		panic("WorkflowMock.StatusFunc: method is nil but Workflow.Status was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		WorkflowName string
	}{
		Ctx:          ctx,
		WorkflowName: workflowName,
	}
	mock.lockStatus.Lock()
	mock.calls.Status = append(mock.calls.Status, callInfo)
	mock.lockStatus.Unlock()
	return mock.StatusFunc(ctx, workflowName)
}

// StatusCalls gets all the calls that were made to Status.
// Check the length with:
//     len(mockedWorkflow.StatusCalls())
func (mock *WorkflowMock) StatusCalls() []struct {
	Ctx          context.Context
	WorkflowName string
} {
	var calls []struct {
		Ctx          context.Context
		WorkflowName string
	}
	mock.lockStatus.RLock()
	calls = mock.calls.Status
	mock.lockStatus.RUnlock()
	return calls
}

// Submit calls SubmitFunc.
func (mock *WorkflowMock) Submit(ctx context.Context, from string, parameters map[string]string, labels map[string]string) (string, error) {
	if mock.SubmitFunc == nil {
		panic("WorkflowMock.SubmitFunc: method is nil but Workflow.Submit was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		From       string
		Parameters map[string]string
		Labels     map[string]string
	}{
		Ctx:        ctx,
		From:       from,
		Parameters: parameters,
		Labels:     labels,
	}
	mock.lockSubmit.Lock()
	mock.calls.Submit = append(mock.calls.Submit, callInfo)
	mock.lockSubmit.Unlock()
	return mock.SubmitFunc(ctx, from, parameters, labels)
}

// SubmitCalls gets all the calls that were made to Submit.
// Check the length with:
//     len(mockedWorkflow.SubmitCalls())
func (mock *WorkflowMock) SubmitCalls() []struct {
	Ctx        context.Context
	From       string
	Parameters map[string]string
	Labels     map[string]string
} {
	var calls []struct {
		Ctx        context.Context
		From       string
		Parameters map[string]string
		Labels     map[string]string
	}
	mock.lockSubmit.RLock()
	calls = mock.calls.Submit
	mock.lockSubmit.RUnlock()
	return calls
}
