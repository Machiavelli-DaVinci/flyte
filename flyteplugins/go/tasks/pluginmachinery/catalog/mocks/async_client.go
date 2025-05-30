// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	catalog "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/catalog"

	mock "github.com/stretchr/testify/mock"
)

// AsyncClient is an autogenerated mock type for the AsyncClient type
type AsyncClient struct {
	mock.Mock
}

type AsyncClient_Expecter struct {
	mock *mock.Mock
}

func (_m *AsyncClient) EXPECT() *AsyncClient_Expecter {
	return &AsyncClient_Expecter{mock: &_m.Mock}
}

// Download provides a mock function with given fields: ctx, requests
func (_m *AsyncClient) Download(ctx context.Context, requests ...catalog.DownloadRequest) (catalog.DownloadFuture, error) {
	_va := make([]interface{}, len(requests))
	for _i := range requests {
		_va[_i] = requests[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 catalog.DownloadFuture
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...catalog.DownloadRequest) (catalog.DownloadFuture, error)); ok {
		return rf(ctx, requests...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...catalog.DownloadRequest) catalog.DownloadFuture); ok {
		r0 = rf(ctx, requests...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.DownloadFuture)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...catalog.DownloadRequest) error); ok {
		r1 = rf(ctx, requests...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AsyncClient_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type AsyncClient_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - requests ...catalog.DownloadRequest
func (_e *AsyncClient_Expecter) Download(ctx interface{}, requests ...interface{}) *AsyncClient_Download_Call {
	return &AsyncClient_Download_Call{Call: _e.mock.On("Download",
		append([]interface{}{ctx}, requests...)...)}
}

func (_c *AsyncClient_Download_Call) Run(run func(ctx context.Context, requests ...catalog.DownloadRequest)) *AsyncClient_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]catalog.DownloadRequest, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(catalog.DownloadRequest)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *AsyncClient_Download_Call) Return(outputFuture catalog.DownloadFuture, err error) *AsyncClient_Download_Call {
	_c.Call.Return(outputFuture, err)
	return _c
}

func (_c *AsyncClient_Download_Call) RunAndReturn(run func(context.Context, ...catalog.DownloadRequest) (catalog.DownloadFuture, error)) *AsyncClient_Download_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx, requests
func (_m *AsyncClient) Upload(ctx context.Context, requests ...catalog.UploadRequest) (catalog.UploadFuture, error) {
	_va := make([]interface{}, len(requests))
	for _i := range requests {
		_va[_i] = requests[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 catalog.UploadFuture
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...catalog.UploadRequest) (catalog.UploadFuture, error)); ok {
		return rf(ctx, requests...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...catalog.UploadRequest) catalog.UploadFuture); ok {
		r0 = rf(ctx, requests...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.UploadFuture)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...catalog.UploadRequest) error); ok {
		r1 = rf(ctx, requests...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AsyncClient_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type AsyncClient_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - requests ...catalog.UploadRequest
func (_e *AsyncClient_Expecter) Upload(ctx interface{}, requests ...interface{}) *AsyncClient_Upload_Call {
	return &AsyncClient_Upload_Call{Call: _e.mock.On("Upload",
		append([]interface{}{ctx}, requests...)...)}
}

func (_c *AsyncClient_Upload_Call) Run(run func(ctx context.Context, requests ...catalog.UploadRequest)) *AsyncClient_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]catalog.UploadRequest, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(catalog.UploadRequest)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *AsyncClient_Upload_Call) Return(putFuture catalog.UploadFuture, err error) *AsyncClient_Upload_Call {
	_c.Call.Return(putFuture, err)
	return _c
}

func (_c *AsyncClient_Upload_Call) RunAndReturn(run func(context.Context, ...catalog.UploadRequest) (catalog.UploadFuture, error)) *AsyncClient_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewAsyncClient creates a new instance of AsyncClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAsyncClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AsyncClient {
	mock := &AsyncClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
