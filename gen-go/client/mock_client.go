// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	models "github.com/Clever/workflow-manager/gen-go/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method
func (m *MockClient) HealthCheck(ctx context.Context) error {
	ret := m.ctrl.Call(m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockClientMockRecorder) HealthCheck(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockClient)(nil).HealthCheck), ctx)
}

// PostStateResource mocks base method
func (m *MockClient) PostStateResource(ctx context.Context, i *models.NewStateResource) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "PostStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStateResource indicates an expected call of PostStateResource
func (mr *MockClientMockRecorder) PostStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStateResource", reflect.TypeOf((*MockClient)(nil).PostStateResource), ctx, i)
}

// DeleteStateResource mocks base method
func (m *MockClient) DeleteStateResource(ctx context.Context, i *models.DeleteStateResourceInput) error {
	ret := m.ctrl.Call(m, "DeleteStateResource", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStateResource indicates an expected call of DeleteStateResource
func (mr *MockClientMockRecorder) DeleteStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStateResource", reflect.TypeOf((*MockClient)(nil).DeleteStateResource), ctx, i)
}

// GetStateResource mocks base method
func (m *MockClient) GetStateResource(ctx context.Context, i *models.GetStateResourceInput) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "GetStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateResource indicates an expected call of GetStateResource
func (mr *MockClientMockRecorder) GetStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateResource", reflect.TypeOf((*MockClient)(nil).GetStateResource), ctx, i)
}

// PutStateResource mocks base method
func (m *MockClient) PutStateResource(ctx context.Context, i *models.PutStateResourceInput) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "PutStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutStateResource indicates an expected call of PutStateResource
func (mr *MockClientMockRecorder) PutStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutStateResource", reflect.TypeOf((*MockClient)(nil).PutStateResource), ctx, i)
}

// GetWorkflowDefinitions mocks base method
func (m *MockClient) GetWorkflowDefinitions(ctx context.Context) ([]models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitions", ctx)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitions indicates an expected call of GetWorkflowDefinitions
func (mr *MockClientMockRecorder) GetWorkflowDefinitions(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitions", reflect.TypeOf((*MockClient)(nil).GetWorkflowDefinitions), ctx)
}

// NewWorkflowDefinition mocks base method
func (m *MockClient) NewWorkflowDefinition(ctx context.Context, i *models.NewWorkflowDefinitionRequest) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "NewWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWorkflowDefinition indicates an expected call of NewWorkflowDefinition
func (mr *MockClientMockRecorder) NewWorkflowDefinition(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWorkflowDefinition", reflect.TypeOf((*MockClient)(nil).NewWorkflowDefinition), ctx, i)
}

// GetWorkflowDefinitionVersionsByName mocks base method
func (m *MockClient) GetWorkflowDefinitionVersionsByName(ctx context.Context, i *models.GetWorkflowDefinitionVersionsByNameInput) ([]models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitionVersionsByName", ctx, i)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitionVersionsByName indicates an expected call of GetWorkflowDefinitionVersionsByName
func (mr *MockClientMockRecorder) GetWorkflowDefinitionVersionsByName(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitionVersionsByName", reflect.TypeOf((*MockClient)(nil).GetWorkflowDefinitionVersionsByName), ctx, i)
}

// UpdateWorkflowDefinition mocks base method
func (m *MockClient) UpdateWorkflowDefinition(ctx context.Context, i *models.UpdateWorkflowDefinitionInput) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "UpdateWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflowDefinition indicates an expected call of UpdateWorkflowDefinition
func (mr *MockClientMockRecorder) UpdateWorkflowDefinition(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowDefinition", reflect.TypeOf((*MockClient)(nil).UpdateWorkflowDefinition), ctx, i)
}

// GetWorkflowDefinitionByNameAndVersion mocks base method
func (m *MockClient) GetWorkflowDefinitionByNameAndVersion(ctx context.Context, i *models.GetWorkflowDefinitionByNameAndVersionInput) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitionByNameAndVersion", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitionByNameAndVersion indicates an expected call of GetWorkflowDefinitionByNameAndVersion
func (mr *MockClientMockRecorder) GetWorkflowDefinitionByNameAndVersion(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitionByNameAndVersion", reflect.TypeOf((*MockClient)(nil).GetWorkflowDefinitionByNameAndVersion), ctx, i)
}

// GetWorkflows mocks base method
func (m *MockClient) GetWorkflows(ctx context.Context, i *models.GetWorkflowsInput) ([]models.Workflow, error) {
	ret := m.ctrl.Call(m, "GetWorkflows", ctx, i)
	ret0, _ := ret[0].([]models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflows indicates an expected call of GetWorkflows
func (mr *MockClientMockRecorder) GetWorkflows(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflows", reflect.TypeOf((*MockClient)(nil).GetWorkflows), ctx, i)
}

// NewGetWorkflowsIter mocks base method
func (m *MockClient) NewGetWorkflowsIter(ctx context.Context, i *models.GetWorkflowsInput) (GetWorkflowsIter, error) {
	ret := m.ctrl.Call(m, "NewGetWorkflowsIter", ctx, i)
	ret0, _ := ret[0].(GetWorkflowsIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGetWorkflowsIter indicates an expected call of NewGetWorkflowsIter
func (mr *MockClientMockRecorder) NewGetWorkflowsIter(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetWorkflowsIter", reflect.TypeOf((*MockClient)(nil).NewGetWorkflowsIter), ctx, i)
}

// StartWorkflow mocks base method
func (m *MockClient) StartWorkflow(ctx context.Context, i *models.StartWorkflowRequest) (*models.Workflow, error) {
	ret := m.ctrl.Call(m, "StartWorkflow", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflow indicates an expected call of StartWorkflow
func (mr *MockClientMockRecorder) StartWorkflow(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflow", reflect.TypeOf((*MockClient)(nil).StartWorkflow), ctx, i)
}

// CancelWorkflow mocks base method
func (m *MockClient) CancelWorkflow(ctx context.Context, i *models.CancelWorkflowInput) error {
	ret := m.ctrl.Call(m, "CancelWorkflow", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelWorkflow indicates an expected call of CancelWorkflow
func (mr *MockClientMockRecorder) CancelWorkflow(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWorkflow", reflect.TypeOf((*MockClient)(nil).CancelWorkflow), ctx, i)
}

// GetWorkflowByID mocks base method
func (m *MockClient) GetWorkflowByID(ctx context.Context, workflowID string) (*models.Workflow, error) {
	ret := m.ctrl.Call(m, "GetWorkflowByID", ctx, workflowID)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByID indicates an expected call of GetWorkflowByID
func (mr *MockClientMockRecorder) GetWorkflowByID(ctx, workflowID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowByID", reflect.TypeOf((*MockClient)(nil).GetWorkflowByID), ctx, workflowID)
}

// ResumeWorkflowByID mocks base method
func (m *MockClient) ResumeWorkflowByID(ctx context.Context, i *models.ResumeWorkflowByIDInput) (*models.Workflow, error) {
	ret := m.ctrl.Call(m, "ResumeWorkflowByID", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResumeWorkflowByID indicates an expected call of ResumeWorkflowByID
func (mr *MockClientMockRecorder) ResumeWorkflowByID(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeWorkflowByID", reflect.TypeOf((*MockClient)(nil).ResumeWorkflowByID), ctx, i)
}

// ResolveWorkflowByID mocks base method
func (m *MockClient) ResolveWorkflowByID(ctx context.Context, workflowID string) error {
	ret := m.ctrl.Call(m, "ResolveWorkflowByID", ctx, workflowID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveWorkflowByID indicates an expected call of ResolveWorkflowByID
func (mr *MockClientMockRecorder) ResolveWorkflowByID(ctx, workflowID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWorkflowByID", reflect.TypeOf((*MockClient)(nil).ResolveWorkflowByID), ctx, workflowID)
}

// MockGetWorkflowsIter is a mock of GetWorkflowsIter interface
type MockGetWorkflowsIter struct {
	ctrl     *gomock.Controller
	recorder *MockGetWorkflowsIterMockRecorder
}

// MockGetWorkflowsIterMockRecorder is the mock recorder for MockGetWorkflowsIter
type MockGetWorkflowsIterMockRecorder struct {
	mock *MockGetWorkflowsIter
}

// NewMockGetWorkflowsIter creates a new mock instance
func NewMockGetWorkflowsIter(ctrl *gomock.Controller) *MockGetWorkflowsIter {
	mock := &MockGetWorkflowsIter{ctrl: ctrl}
	mock.recorder = &MockGetWorkflowsIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetWorkflowsIter) EXPECT() *MockGetWorkflowsIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockGetWorkflowsIter) Next(arg0 *models.Workflow) bool {
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockGetWorkflowsIterMockRecorder) Next(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockGetWorkflowsIter)(nil).Next), arg0)
}

// Err mocks base method
func (m *MockGetWorkflowsIter) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockGetWorkflowsIterMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockGetWorkflowsIter)(nil).Err))
}
