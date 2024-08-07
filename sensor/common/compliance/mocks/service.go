// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/service.go -source service.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	central "github.com/stackrox/rox/generated/internalapi/central"
	compliance "github.com/stackrox/rox/generated/internalapi/compliance"
	sensor "github.com/stackrox/rox/generated/internalapi/sensor"
	storage "github.com/stackrox/rox/generated/storage"
	centralsensor "github.com/stackrox/rox/pkg/centralsensor"
	common "github.com/stackrox/rox/sensor/common"
	message "github.com/stackrox/rox/sensor/common/message"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AuditEvents mocks base method.
func (m *MockService) AuditEvents() chan *sensor.AuditEvents {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuditEvents")
	ret0, _ := ret[0].(chan *sensor.AuditEvents)
	return ret0
}

// AuditEvents indicates an expected call of AuditEvents.
func (mr *MockServiceMockRecorder) AuditEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditEvents", reflect.TypeOf((*MockService)(nil).AuditEvents))
}

// AuthFuncOverride mocks base method.
func (m *MockService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthFuncOverride", ctx, fullMethodName)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthFuncOverride indicates an expected call of AuthFuncOverride.
func (mr *MockServiceMockRecorder) AuthFuncOverride(ctx, fullMethodName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthFuncOverride", reflect.TypeOf((*MockService)(nil).AuthFuncOverride), ctx, fullMethodName)
}

// Capabilities mocks base method.
func (m *MockService) Capabilities() []centralsensor.SensorCapability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].([]centralsensor.SensorCapability)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockServiceMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockService)(nil).Capabilities))
}

// Communicate mocks base method.
func (m *MockService) Communicate(arg0 grpc.BidiStreamingServer[sensor.MsgFromCompliance, sensor.MsgToCompliance]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Communicate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Communicate indicates an expected call of Communicate.
func (mr *MockServiceMockRecorder) Communicate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Communicate", reflect.TypeOf((*MockService)(nil).Communicate), arg0)
}

// NodeInventories mocks base method.
func (m *MockService) NodeInventories() <-chan *storage.NodeInventory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeInventories")
	ret0, _ := ret[0].(<-chan *storage.NodeInventory)
	return ret0
}

// NodeInventories indicates an expected call of NodeInventories.
func (mr *MockServiceMockRecorder) NodeInventories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeInventories", reflect.TypeOf((*MockService)(nil).NodeInventories))
}

// Notify mocks base method.
func (m *MockService) Notify(e common.SensorComponentEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", e)
}

// Notify indicates an expected call of Notify.
func (mr *MockServiceMockRecorder) Notify(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockService)(nil).Notify), e)
}

// Output mocks base method.
func (m *MockService) Output() chan *compliance.ComplianceReturn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(chan *compliance.ComplianceReturn)
	return ret0
}

// Output indicates an expected call of Output.
func (mr *MockServiceMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockService)(nil).Output))
}

// ProcessMessage mocks base method.
func (m *MockService) ProcessMessage(msg *central.MsgToSensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockServiceMockRecorder) ProcessMessage(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockService)(nil).ProcessMessage), msg)
}

// RegisterServiceHandler mocks base method.
func (m *MockService) RegisterServiceHandler(arg0 context.Context, arg1 *runtime.ServeMux, arg2 *grpc.ClientConn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterServiceHandler", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterServiceHandler indicates an expected call of RegisterServiceHandler.
func (mr *MockServiceMockRecorder) RegisterServiceHandler(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServiceHandler", reflect.TypeOf((*MockService)(nil).RegisterServiceHandler), arg0, arg1, arg2)
}

// RegisterServiceServer mocks base method.
func (m *MockService) RegisterServiceServer(server *grpc.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterServiceServer", server)
}

// RegisterServiceServer indicates an expected call of RegisterServiceServer.
func (mr *MockServiceMockRecorder) RegisterServiceServer(server any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServiceServer", reflect.TypeOf((*MockService)(nil).RegisterServiceServer), server)
}

// ResponsesC mocks base method.
func (m *MockService) ResponsesC() <-chan *message.ExpiringMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponsesC")
	ret0, _ := ret[0].(<-chan *message.ExpiringMessage)
	return ret0
}

// ResponsesC indicates an expected call of ResponsesC.
func (mr *MockServiceMockRecorder) ResponsesC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponsesC", reflect.TypeOf((*MockService)(nil).ResponsesC))
}

// RunScrape mocks base method.
func (m *MockService) RunScrape(msg *sensor.MsgToCompliance) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunScrape", msg)
	ret0, _ := ret[0].(int)
	return ret0
}

// RunScrape indicates an expected call of RunScrape.
func (mr *MockServiceMockRecorder) RunScrape(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunScrape", reflect.TypeOf((*MockService)(nil).RunScrape), msg)
}

// Start mocks base method.
func (m *MockService) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockServiceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockService)(nil).Start))
}

// Stop mocks base method.
func (m *MockService) Stop(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", err)
}

// Stop indicates an expected call of Stop.
func (mr *MockServiceMockRecorder) Stop(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockService)(nil).Stop), err)
}
