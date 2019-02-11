// Code generated by MockGen. DO NOT EDIT.
// Source: spider_service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	models "github.com/getumen/gogo_crawler/domains/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSpiderService is a mock of SpiderService interface
type MockSpiderService struct {
	ctrl     *gomock.Controller
	recorder *MockSpiderServiceMockRecorder
}

// MockSpiderServiceMockRecorder is the mock recorder for MockSpiderService
type MockSpiderServiceMockRecorder struct {
	mock *MockSpiderService
}

// NewMockSpiderService creates a new mock instance
func NewMockSpiderService(ctrl *gomock.Controller) *MockSpiderService {
	mock := &MockSpiderService{ctrl: ctrl}
	mock.recorder = &MockSpiderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpiderService) EXPECT() *MockSpiderServiceMockRecorder {
	return m.recorder
}

// ParseResponse mocks base method
func (m *MockSpiderService) ParseResponse(ctx context.Context, allowedDomainRegexp string, in <-chan *models.Response, out chan<- *models.Request) {
	m.ctrl.Call(m, "ParseResponse", ctx, allowedDomainRegexp, in, out)
}

// ParseResponse indicates an expected call of ParseResponse
func (mr *MockSpiderServiceMockRecorder) ParseResponse(ctx, allowedDomainRegexp, in, out interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseResponse", reflect.TypeOf((*MockSpiderService)(nil).ParseResponse), ctx, allowedDomainRegexp, in, out)
}
