// Code generated by MockGen. DO NOT EDIT.
// Source: cluster_client.go
//
// Generated by this command:
//
//	mockgen -source=cluster_client.go -package=common -destination=mock_clusterclient.go
//
// Package common is a generated GoMock package.
package common

import (
	context "context"
	reflect "reflect"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterClient is a mock of ClusterClient interface.
type MockClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientMockRecorder
}

// MockClusterClientMockRecorder is the mock recorder for MockClusterClient.
type MockClusterClientMockRecorder struct {
	mock *MockClusterClient
}

// NewMockClusterClient creates a new mock instance.
func NewMockClusterClient(ctrl *gomock.Controller) *MockClusterClient {
	mock := &MockClusterClient{ctrl: ctrl}
	mock.recorder = &MockClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterClient) EXPECT() *MockClusterClientMockRecorder {
	return m.recorder
}

// FetchCluster mocks base method.
func (m *MockClusterClient) FetchCluster(ctx context.Context, clusterId string) (*v1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCluster", ctx, clusterId)
	ret0, _ := ret[0].(*v1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCluster indicates an expected call of FetchCluster.
func (mr *MockClusterClientMockRecorder) FetchCluster(ctx, clusterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCluster", reflect.TypeOf((*MockClusterClient)(nil).FetchCluster), ctx, clusterId)
}
