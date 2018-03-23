package servicediscovery

import (
	gomock "github.com/golang/mock/gomock"
	dns "github.com/miekg/dns"
	time "time"
)

// Mock of DnsClient interface
type MockDnsClient struct {
	ctrl     *gomock.Controller
	recorder *_MockDnsClientRecorder
}

// Recorder for MockDnsClient (not exported)
type _MockDnsClientRecorder struct {
	mock *MockDnsClient
}

func NewMockDnsClient(ctrl *gomock.Controller) *MockDnsClient {
	mock := &MockDnsClient{ctrl: ctrl}
	mock.recorder = &_MockDnsClientRecorder{mock}
	return mock
}

func (_m *MockDnsClient) EXPECT() *_MockDnsClientRecorder {
	return _m.recorder
}

func (_m *MockDnsClient) Exchange(_param0 *dns.Msg, _param1 string) (*dns.Msg, time.Duration, error) {
	ret := _m.ctrl.Call(_m, "Exchange", _param0, _param1)
	ret0, _ := ret[0].(*dns.Msg)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockDnsClientRecorder) Exchange(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exchange", arg0, arg1)
}

// Mock of ServiceDiscovery interface
type MockServiceDiscovery struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceDiscoveryRecorder
}

// Recorder for MockServiceDiscovery (not exported)
type _MockServiceDiscoveryRecorder struct {
	mock *MockServiceDiscovery
}

func NewMockServiceDiscovery(ctrl *gomock.Controller) *MockServiceDiscovery {
	mock := &MockServiceDiscovery{ctrl: ctrl}
	mock.recorder = &_MockServiceDiscoveryRecorder{mock}
	return mock
}

func (_m *MockServiceDiscovery) EXPECT() *_MockServiceDiscoveryRecorder {
	return _m.recorder
}

func (_m *MockServiceDiscovery) DiscoverService(serviceName string) (string, string, error) {
	ret := _m.ctrl.Call(_m, "DiscoverService", serviceName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockServiceDiscoveryRecorder) DiscoverService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscoverService", arg0)
}

func (_m *MockServiceDiscovery) DiscoverAllServiceInstances(serviceName string) ([]ServiceInstance, error) {
	ret := _m.ctrl.Call(_m, "DiscoverAllServiceInstances", serviceName)
	ret0, _ := ret[0].([]ServiceInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceDiscoveryRecorder) DiscoverAllServiceInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscoverAllServiceInstances", arg0)
}

