// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -package engine -source interface.go -destination interface_mock.go
//

// Package engine is a generated GoMock package.
package engine

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	types "github.com/uber/cadence/common/types"
	common "github.com/uber/cadence/service/history/common"
	events "github.com/uber/cadence/service/history/events"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
	isgomock struct{}
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// CountDLQMessages mocks base method.
func (m *MockEngine) CountDLQMessages(ctx context.Context, forceFetch bool) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDLQMessages", ctx, forceFetch)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDLQMessages indicates an expected call of CountDLQMessages.
func (mr *MockEngineMockRecorder) CountDLQMessages(ctx, forceFetch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDLQMessages", reflect.TypeOf((*MockEngine)(nil).CountDLQMessages), ctx, forceFetch)
}

// DescribeMutableState mocks base method.
func (m *MockEngine) DescribeMutableState(ctx context.Context, request *types.DescribeMutableStateRequest) (*types.DescribeMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMutableState", ctx, request)
	ret0, _ := ret[0].(*types.DescribeMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMutableState indicates an expected call of DescribeMutableState.
func (mr *MockEngineMockRecorder) DescribeMutableState(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMutableState", reflect.TypeOf((*MockEngine)(nil).DescribeMutableState), ctx, request)
}

// DescribeTimerQueue mocks base method.
func (m *MockEngine) DescribeTimerQueue(ctx context.Context, clusterName string) (*types.DescribeQueueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTimerQueue", ctx, clusterName)
	ret0, _ := ret[0].(*types.DescribeQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTimerQueue indicates an expected call of DescribeTimerQueue.
func (mr *MockEngineMockRecorder) DescribeTimerQueue(ctx, clusterName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTimerQueue", reflect.TypeOf((*MockEngine)(nil).DescribeTimerQueue), ctx, clusterName)
}

// DescribeTransferQueue mocks base method.
func (m *MockEngine) DescribeTransferQueue(ctx context.Context, clusterName string) (*types.DescribeQueueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTransferQueue", ctx, clusterName)
	ret0, _ := ret[0].(*types.DescribeQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTransferQueue indicates an expected call of DescribeTransferQueue.
func (mr *MockEngineMockRecorder) DescribeTransferQueue(ctx, clusterName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransferQueue", reflect.TypeOf((*MockEngine)(nil).DescribeTransferQueue), ctx, clusterName)
}

// DescribeWorkflowExecution mocks base method.
func (m *MockEngine) DescribeWorkflowExecution(ctx context.Context, request *types.HistoryDescribeWorkflowExecutionRequest) (*types.DescribeWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*types.DescribeWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflowExecution indicates an expected call of DescribeWorkflowExecution.
func (mr *MockEngineMockRecorder) DescribeWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).DescribeWorkflowExecution), ctx, request)
}

// GetDLQReplicationMessages mocks base method.
func (m *MockEngine) GetDLQReplicationMessages(ctx context.Context, taskInfos []*types.ReplicationTaskInfo) ([]*types.ReplicationTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLQReplicationMessages", ctx, taskInfos)
	ret0, _ := ret[0].([]*types.ReplicationTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQReplicationMessages indicates an expected call of GetDLQReplicationMessages.
func (mr *MockEngineMockRecorder) GetDLQReplicationMessages(ctx, taskInfos any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQReplicationMessages", reflect.TypeOf((*MockEngine)(nil).GetDLQReplicationMessages), ctx, taskInfos)
}

// GetMutableState mocks base method.
func (m *MockEngine) GetMutableState(ctx context.Context, request *types.GetMutableStateRequest) (*types.GetMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMutableState", ctx, request)
	ret0, _ := ret[0].(*types.GetMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMutableState indicates an expected call of GetMutableState.
func (mr *MockEngineMockRecorder) GetMutableState(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMutableState", reflect.TypeOf((*MockEngine)(nil).GetMutableState), ctx, request)
}

// GetReplicationMessages mocks base method.
func (m *MockEngine) GetReplicationMessages(ctx context.Context, pollingCluster string, lastReadMessageID int64) (*types.ReplicationMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationMessages", ctx, pollingCluster, lastReadMessageID)
	ret0, _ := ret[0].(*types.ReplicationMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicationMessages indicates an expected call of GetReplicationMessages.
func (mr *MockEngineMockRecorder) GetReplicationMessages(ctx, pollingCluster, lastReadMessageID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationMessages", reflect.TypeOf((*MockEngine)(nil).GetReplicationMessages), ctx, pollingCluster, lastReadMessageID)
}

// MergeDLQMessages mocks base method.
func (m *MockEngine) MergeDLQMessages(ctx context.Context, messagesRequest *types.MergeDLQMessagesRequest) (*types.MergeDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(*types.MergeDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeDLQMessages indicates an expected call of MergeDLQMessages.
func (mr *MockEngineMockRecorder) MergeDLQMessages(ctx, messagesRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeDLQMessages", reflect.TypeOf((*MockEngine)(nil).MergeDLQMessages), ctx, messagesRequest)
}

// NotifyNewHistoryEvent mocks base method.
func (m *MockEngine) NotifyNewHistoryEvent(event *events.Notification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewHistoryEvent", event)
}

// NotifyNewHistoryEvent indicates an expected call of NotifyNewHistoryEvent.
func (mr *MockEngineMockRecorder) NotifyNewHistoryEvent(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewHistoryEvent", reflect.TypeOf((*MockEngine)(nil).NotifyNewHistoryEvent), event)
}

// NotifyNewReplicationTasks mocks base method.
func (m *MockEngine) NotifyNewReplicationTasks(info *common.NotifyTaskInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewReplicationTasks", info)
}

// NotifyNewReplicationTasks indicates an expected call of NotifyNewReplicationTasks.
func (mr *MockEngineMockRecorder) NotifyNewReplicationTasks(info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewReplicationTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewReplicationTasks), info)
}

// NotifyNewTimerTasks mocks base method.
func (m *MockEngine) NotifyNewTimerTasks(info *common.NotifyTaskInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTimerTasks", info)
}

// NotifyNewTimerTasks indicates an expected call of NotifyNewTimerTasks.
func (mr *MockEngineMockRecorder) NotifyNewTimerTasks(info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTimerTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewTimerTasks), info)
}

// NotifyNewTransferTasks mocks base method.
func (m *MockEngine) NotifyNewTransferTasks(info *common.NotifyTaskInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTransferTasks", info)
}

// NotifyNewTransferTasks indicates an expected call of NotifyNewTransferTasks.
func (mr *MockEngineMockRecorder) NotifyNewTransferTasks(info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTransferTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewTransferTasks), info)
}

// PollMutableState mocks base method.
func (m *MockEngine) PollMutableState(ctx context.Context, request *types.PollMutableStateRequest) (*types.PollMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollMutableState", ctx, request)
	ret0, _ := ret[0].(*types.PollMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollMutableState indicates an expected call of PollMutableState.
func (mr *MockEngineMockRecorder) PollMutableState(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollMutableState", reflect.TypeOf((*MockEngine)(nil).PollMutableState), ctx, request)
}

// PurgeDLQMessages mocks base method.
func (m *MockEngine) PurgeDLQMessages(ctx context.Context, messagesRequest *types.PurgeDLQMessagesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeDLQMessages indicates an expected call of PurgeDLQMessages.
func (mr *MockEngineMockRecorder) PurgeDLQMessages(ctx, messagesRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDLQMessages", reflect.TypeOf((*MockEngine)(nil).PurgeDLQMessages), ctx, messagesRequest)
}

// QueryWorkflow mocks base method.
func (m *MockEngine) QueryWorkflow(ctx context.Context, request *types.HistoryQueryWorkflowRequest) (*types.HistoryQueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryWorkflow", ctx, request)
	ret0, _ := ret[0].(*types.HistoryQueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow.
func (mr *MockEngineMockRecorder) QueryWorkflow(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockEngine)(nil).QueryWorkflow), ctx, request)
}

// ReadDLQMessages mocks base method.
func (m *MockEngine) ReadDLQMessages(ctx context.Context, messagesRequest *types.ReadDLQMessagesRequest) (*types.ReadDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(*types.ReadDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDLQMessages indicates an expected call of ReadDLQMessages.
func (mr *MockEngineMockRecorder) ReadDLQMessages(ctx, messagesRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDLQMessages", reflect.TypeOf((*MockEngine)(nil).ReadDLQMessages), ctx, messagesRequest)
}

// ReapplyEvents mocks base method.
func (m *MockEngine) ReapplyEvents(ctx context.Context, domainUUID, workflowID, runID string, events []*types.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReapplyEvents", ctx, domainUUID, workflowID, runID, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReapplyEvents indicates an expected call of ReapplyEvents.
func (mr *MockEngineMockRecorder) ReapplyEvents(ctx, domainUUID, workflowID, runID, events any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapplyEvents", reflect.TypeOf((*MockEngine)(nil).ReapplyEvents), ctx, domainUUID, workflowID, runID, events)
}

// RecordActivityTaskHeartbeat mocks base method.
func (m *MockEngine) RecordActivityTaskHeartbeat(ctx context.Context, request *types.HistoryRecordActivityTaskHeartbeatRequest) (*types.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", ctx, request)
	ret0, _ := ret[0].(*types.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeat indicates an expected call of RecordActivityTaskHeartbeat.
func (mr *MockEngineMockRecorder) RecordActivityTaskHeartbeat(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeat", reflect.TypeOf((*MockEngine)(nil).RecordActivityTaskHeartbeat), ctx, request)
}

// RecordActivityTaskStarted mocks base method.
func (m *MockEngine) RecordActivityTaskStarted(ctx context.Context, request *types.RecordActivityTaskStartedRequest) (*types.RecordActivityTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskStarted", ctx, request)
	ret0, _ := ret[0].(*types.RecordActivityTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskStarted indicates an expected call of RecordActivityTaskStarted.
func (mr *MockEngineMockRecorder) RecordActivityTaskStarted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskStarted", reflect.TypeOf((*MockEngine)(nil).RecordActivityTaskStarted), ctx, request)
}

// RecordChildExecutionCompleted mocks base method.
func (m *MockEngine) RecordChildExecutionCompleted(ctx context.Context, request *types.RecordChildExecutionCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordChildExecutionCompleted", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordChildExecutionCompleted indicates an expected call of RecordChildExecutionCompleted.
func (mr *MockEngineMockRecorder) RecordChildExecutionCompleted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordChildExecutionCompleted", reflect.TypeOf((*MockEngine)(nil).RecordChildExecutionCompleted), ctx, request)
}

// RecordDecisionTaskStarted mocks base method.
func (m *MockEngine) RecordDecisionTaskStarted(ctx context.Context, request *types.RecordDecisionTaskStartedRequest) (*types.RecordDecisionTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordDecisionTaskStarted", ctx, request)
	ret0, _ := ret[0].(*types.RecordDecisionTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordDecisionTaskStarted indicates an expected call of RecordDecisionTaskStarted.
func (mr *MockEngineMockRecorder) RecordDecisionTaskStarted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordDecisionTaskStarted", reflect.TypeOf((*MockEngine)(nil).RecordDecisionTaskStarted), ctx, request)
}

// RefreshWorkflowTasks mocks base method.
func (m *MockEngine) RefreshWorkflowTasks(ctx context.Context, domainUUID string, execution types.WorkflowExecution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWorkflowTasks", ctx, domainUUID, execution)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWorkflowTasks indicates an expected call of RefreshWorkflowTasks.
func (mr *MockEngineMockRecorder) RefreshWorkflowTasks(ctx, domainUUID, execution any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWorkflowTasks", reflect.TypeOf((*MockEngine)(nil).RefreshWorkflowTasks), ctx, domainUUID, execution)
}

// RemoveSignalMutableState mocks base method.
func (m *MockEngine) RemoveSignalMutableState(ctx context.Context, request *types.RemoveSignalMutableStateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSignalMutableState", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSignalMutableState indicates an expected call of RemoveSignalMutableState.
func (mr *MockEngineMockRecorder) RemoveSignalMutableState(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSignalMutableState", reflect.TypeOf((*MockEngine)(nil).RemoveSignalMutableState), ctx, request)
}

// ReplicateEventsV2 mocks base method.
func (m *MockEngine) ReplicateEventsV2(ctx context.Context, request *types.ReplicateEventsV2Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateEventsV2", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateEventsV2 indicates an expected call of ReplicateEventsV2.
func (mr *MockEngineMockRecorder) ReplicateEventsV2(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateEventsV2", reflect.TypeOf((*MockEngine)(nil).ReplicateEventsV2), ctx, request)
}

// RequestCancelWorkflowExecution mocks base method.
func (m *MockEngine) RequestCancelWorkflowExecution(ctx context.Context, request *types.HistoryRequestCancelWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestCancelWorkflowExecution indicates an expected call of RequestCancelWorkflowExecution.
func (mr *MockEngineMockRecorder) RequestCancelWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCancelWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).RequestCancelWorkflowExecution), ctx, request)
}

// ResetStickyTaskList mocks base method.
func (m *MockEngine) ResetStickyTaskList(ctx context.Context, resetRequest *types.HistoryResetStickyTaskListRequest) (*types.HistoryResetStickyTaskListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStickyTaskList", ctx, resetRequest)
	ret0, _ := ret[0].(*types.HistoryResetStickyTaskListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetStickyTaskList indicates an expected call of ResetStickyTaskList.
func (mr *MockEngineMockRecorder) ResetStickyTaskList(ctx, resetRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStickyTaskList", reflect.TypeOf((*MockEngine)(nil).ResetStickyTaskList), ctx, resetRequest)
}

// ResetTimerQueue mocks base method.
func (m *MockEngine) ResetTimerQueue(ctx context.Context, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetTimerQueue", ctx, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetTimerQueue indicates an expected call of ResetTimerQueue.
func (mr *MockEngineMockRecorder) ResetTimerQueue(ctx, clusterName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTimerQueue", reflect.TypeOf((*MockEngine)(nil).ResetTimerQueue), ctx, clusterName)
}

// ResetTransferQueue mocks base method.
func (m *MockEngine) ResetTransferQueue(ctx context.Context, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetTransferQueue", ctx, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetTransferQueue indicates an expected call of ResetTransferQueue.
func (mr *MockEngineMockRecorder) ResetTransferQueue(ctx, clusterName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTransferQueue", reflect.TypeOf((*MockEngine)(nil).ResetTransferQueue), ctx, clusterName)
}

// ResetWorkflowExecution mocks base method.
func (m *MockEngine) ResetWorkflowExecution(ctx context.Context, request *types.HistoryResetWorkflowExecutionRequest) (*types.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*types.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution.
func (mr *MockEngineMockRecorder) ResetWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).ResetWorkflowExecution), ctx, request)
}

// RespondActivityTaskCanceled mocks base method.
func (m *MockEngine) RespondActivityTaskCanceled(ctx context.Context, request *types.HistoryRespondActivityTaskCanceledRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceled indicates an expected call of RespondActivityTaskCanceled.
func (mr *MockEngineMockRecorder) RespondActivityTaskCanceled(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceled", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskCanceled), ctx, request)
}

// RespondActivityTaskCompleted mocks base method.
func (m *MockEngine) RespondActivityTaskCompleted(ctx context.Context, request *types.HistoryRespondActivityTaskCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompleted indicates an expected call of RespondActivityTaskCompleted.
func (mr *MockEngineMockRecorder) RespondActivityTaskCompleted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompleted", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskCompleted), ctx, request)
}

// RespondActivityTaskFailed mocks base method.
func (m *MockEngine) RespondActivityTaskFailed(ctx context.Context, request *types.HistoryRespondActivityTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailed indicates an expected call of RespondActivityTaskFailed.
func (mr *MockEngineMockRecorder) RespondActivityTaskFailed(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailed", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskFailed), ctx, request)
}

// RespondDecisionTaskCompleted mocks base method.
func (m *MockEngine) RespondDecisionTaskCompleted(ctx context.Context, request *types.HistoryRespondDecisionTaskCompletedRequest) (*types.HistoryRespondDecisionTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskCompleted", ctx, request)
	ret0, _ := ret[0].(*types.HistoryRespondDecisionTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondDecisionTaskCompleted indicates an expected call of RespondDecisionTaskCompleted.
func (mr *MockEngineMockRecorder) RespondDecisionTaskCompleted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskCompleted", reflect.TypeOf((*MockEngine)(nil).RespondDecisionTaskCompleted), ctx, request)
}

// RespondDecisionTaskFailed mocks base method.
func (m *MockEngine) RespondDecisionTaskFailed(ctx context.Context, request *types.HistoryRespondDecisionTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskFailed", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondDecisionTaskFailed indicates an expected call of RespondDecisionTaskFailed.
func (mr *MockEngineMockRecorder) RespondDecisionTaskFailed(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskFailed", reflect.TypeOf((*MockEngine)(nil).RespondDecisionTaskFailed), ctx, request)
}

// ScheduleDecisionTask mocks base method.
func (m *MockEngine) ScheduleDecisionTask(ctx context.Context, request *types.ScheduleDecisionTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleDecisionTask", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleDecisionTask indicates an expected call of ScheduleDecisionTask.
func (mr *MockEngineMockRecorder) ScheduleDecisionTask(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleDecisionTask", reflect.TypeOf((*MockEngine)(nil).ScheduleDecisionTask), ctx, request)
}

// SignalWithStartWorkflowExecution mocks base method.
func (m *MockEngine) SignalWithStartWorkflowExecution(ctx context.Context, request *types.HistorySignalWithStartWorkflowExecutionRequest) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecution indicates an expected call of SignalWithStartWorkflowExecution.
func (mr *MockEngineMockRecorder) SignalWithStartWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).SignalWithStartWorkflowExecution), ctx, request)
}

// SignalWorkflowExecution mocks base method.
func (m *MockEngine) SignalWorkflowExecution(ctx context.Context, request *types.HistorySignalWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalWorkflowExecution indicates an expected call of SignalWorkflowExecution.
func (mr *MockEngineMockRecorder) SignalWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).SignalWorkflowExecution), ctx, request)
}

// Start mocks base method.
func (m *MockEngine) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockEngineMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start))
}

// StartWorkflowExecution mocks base method.
func (m *MockEngine) StartWorkflowExecution(ctx context.Context, request *types.HistoryStartWorkflowExecutionRequest) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecution indicates an expected call of StartWorkflowExecution.
func (mr *MockEngineMockRecorder) StartWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).StartWorkflowExecution), ctx, request)
}

// Stop mocks base method.
func (m *MockEngine) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockEngineMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEngine)(nil).Stop))
}

// SyncActivity mocks base method.
func (m *MockEngine) SyncActivity(ctx context.Context, request *types.SyncActivityRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncActivity", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncActivity indicates an expected call of SyncActivity.
func (mr *MockEngineMockRecorder) SyncActivity(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncActivity", reflect.TypeOf((*MockEngine)(nil).SyncActivity), ctx, request)
}

// SyncShardStatus mocks base method.
func (m *MockEngine) SyncShardStatus(ctx context.Context, request *types.SyncShardStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncShardStatus", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncShardStatus indicates an expected call of SyncShardStatus.
func (mr *MockEngineMockRecorder) SyncShardStatus(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncShardStatus", reflect.TypeOf((*MockEngine)(nil).SyncShardStatus), ctx, request)
}

// TerminateWorkflowExecution mocks base method.
func (m *MockEngine) TerminateWorkflowExecution(ctx context.Context, request *types.HistoryTerminateWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateWorkflowExecution indicates an expected call of TerminateWorkflowExecution.
func (mr *MockEngineMockRecorder) TerminateWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).TerminateWorkflowExecution), ctx, request)
}
