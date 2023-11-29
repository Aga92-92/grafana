// Code generated by mockery v2.34.2. DO NOT EDIT.

package alertmanager_mock

import (
	context "context"

	definitions "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"

	notifier "github.com/grafana/grafana/pkg/services/ngalert/notifier"

	notify "github.com/grafana/alerting/notify"

	v2models "github.com/prometheus/alertmanager/api/v2/models"
)

// RemoteAlertmanagerMock is an autogenerated mock type for the remoteAlertmanager type
type RemoteAlertmanagerMock struct {
	mock.Mock
}

type RemoteAlertmanagerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *RemoteAlertmanagerMock) EXPECT() *RemoteAlertmanagerMock_Expecter {
	return &RemoteAlertmanagerMock_Expecter{mock: &_m.Mock}
}

// ApplyConfig provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) ApplyConfig(_a0 context.Context, _a1 *models.AlertConfiguration) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.AlertConfiguration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_ApplyConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyConfig'
type RemoteAlertmanagerMock_ApplyConfig_Call struct {
	*mock.Call
}

// ApplyConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.AlertConfiguration
func (_e *RemoteAlertmanagerMock_Expecter) ApplyConfig(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_ApplyConfig_Call {
	return &RemoteAlertmanagerMock_ApplyConfig_Call{Call: _e.mock.On("ApplyConfig", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_ApplyConfig_Call) Run(run func(_a0 context.Context, _a1 *models.AlertConfiguration)) *RemoteAlertmanagerMock_ApplyConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.AlertConfiguration))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_ApplyConfig_Call) Return(_a0 error) *RemoteAlertmanagerMock_ApplyConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_ApplyConfig_Call) RunAndReturn(run func(context.Context, *models.AlertConfiguration) error) *RemoteAlertmanagerMock_ApplyConfig_Call {
	_c.Call.Return(run)
	return _c
}

// CleanUp provides a mock function with given fields:
func (_m *RemoteAlertmanagerMock) CleanUp() {
	_m.Called()
}

// RemoteAlertmanagerMock_CleanUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CleanUp'
type RemoteAlertmanagerMock_CleanUp_Call struct {
	*mock.Call
}

// CleanUp is a helper method to define mock.On call
func (_e *RemoteAlertmanagerMock_Expecter) CleanUp() *RemoteAlertmanagerMock_CleanUp_Call {
	return &RemoteAlertmanagerMock_CleanUp_Call{Call: _e.mock.On("CleanUp")}
}

func (_c *RemoteAlertmanagerMock_CleanUp_Call) Run(run func()) *RemoteAlertmanagerMock_CleanUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_CleanUp_Call) Return() *RemoteAlertmanagerMock_CleanUp_Call {
	_c.Call.Return()
	return _c
}

func (_c *RemoteAlertmanagerMock_CleanUp_Call) RunAndReturn(run func()) *RemoteAlertmanagerMock_CleanUp_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSilence provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) CreateSilence(_a0 context.Context, _a1 *v2models.PostableSilence) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v2models.PostableSilence) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v2models.PostableSilence) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v2models.PostableSilence) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_CreateSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSilence'
type RemoteAlertmanagerMock_CreateSilence_Call struct {
	*mock.Call
}

// CreateSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v2models.PostableSilence
func (_e *RemoteAlertmanagerMock_Expecter) CreateSilence(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_CreateSilence_Call {
	return &RemoteAlertmanagerMock_CreateSilence_Call{Call: _e.mock.On("CreateSilence", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_CreateSilence_Call) Run(run func(_a0 context.Context, _a1 *v2models.PostableSilence)) *RemoteAlertmanagerMock_CreateSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v2models.PostableSilence))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_CreateSilence_Call) Return(_a0 string, _a1 error) *RemoteAlertmanagerMock_CreateSilence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_CreateSilence_Call) RunAndReturn(run func(context.Context, *v2models.PostableSilence) (string, error)) *RemoteAlertmanagerMock_CreateSilence_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSilence provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) DeleteSilence(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_DeleteSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSilence'
type RemoteAlertmanagerMock_DeleteSilence_Call struct {
	*mock.Call
}

// DeleteSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *RemoteAlertmanagerMock_Expecter) DeleteSilence(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_DeleteSilence_Call {
	return &RemoteAlertmanagerMock_DeleteSilence_Call{Call: _e.mock.On("DeleteSilence", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_DeleteSilence_Call) Run(run func(_a0 context.Context, _a1 string)) *RemoteAlertmanagerMock_DeleteSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_DeleteSilence_Call) Return(_a0 error) *RemoteAlertmanagerMock_DeleteSilence_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_DeleteSilence_Call) RunAndReturn(run func(context.Context, string) error) *RemoteAlertmanagerMock_DeleteSilence_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlertGroups provides a mock function with given fields: ctx, active, silenced, inhibited, filter, receiver
func (_m *RemoteAlertmanagerMock) GetAlertGroups(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string) (v2models.AlertGroups, error) {
	ret := _m.Called(ctx, active, silenced, inhibited, filter, receiver)

	var r0 v2models.AlertGroups
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) (v2models.AlertGroups, error)); ok {
		return rf(ctx, active, silenced, inhibited, filter, receiver)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) v2models.AlertGroups); ok {
		r0 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.AlertGroups)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, bool, bool, []string, string) error); ok {
		r1 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_GetAlertGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertGroups'
type RemoteAlertmanagerMock_GetAlertGroups_Call struct {
	*mock.Call
}

// GetAlertGroups is a helper method to define mock.On call
//   - ctx context.Context
//   - active bool
//   - silenced bool
//   - inhibited bool
//   - filter []string
//   - receiver string
func (_e *RemoteAlertmanagerMock_Expecter) GetAlertGroups(ctx interface{}, active interface{}, silenced interface{}, inhibited interface{}, filter interface{}, receiver interface{}) *RemoteAlertmanagerMock_GetAlertGroups_Call {
	return &RemoteAlertmanagerMock_GetAlertGroups_Call{Call: _e.mock.On("GetAlertGroups", ctx, active, silenced, inhibited, filter, receiver)}
}

func (_c *RemoteAlertmanagerMock_GetAlertGroups_Call) Run(run func(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string)) *RemoteAlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(bool), args[3].(bool), args[4].([]string), args[5].(string))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_GetAlertGroups_Call) Return(_a0 v2models.AlertGroups, _a1 error) *RemoteAlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_GetAlertGroups_Call) RunAndReturn(run func(context.Context, bool, bool, bool, []string, string) (v2models.AlertGroups, error)) *RemoteAlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlerts provides a mock function with given fields: ctx, active, silenced, inhibited, filter, receiver
func (_m *RemoteAlertmanagerMock) GetAlerts(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string) (v2models.GettableAlerts, error) {
	ret := _m.Called(ctx, active, silenced, inhibited, filter, receiver)

	var r0 v2models.GettableAlerts
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) (v2models.GettableAlerts, error)); ok {
		return rf(ctx, active, silenced, inhibited, filter, receiver)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) v2models.GettableAlerts); ok {
		r0 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.GettableAlerts)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, bool, bool, []string, string) error); ok {
		r1 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_GetAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlerts'
type RemoteAlertmanagerMock_GetAlerts_Call struct {
	*mock.Call
}

// GetAlerts is a helper method to define mock.On call
//   - ctx context.Context
//   - active bool
//   - silenced bool
//   - inhibited bool
//   - filter []string
//   - receiver string
func (_e *RemoteAlertmanagerMock_Expecter) GetAlerts(ctx interface{}, active interface{}, silenced interface{}, inhibited interface{}, filter interface{}, receiver interface{}) *RemoteAlertmanagerMock_GetAlerts_Call {
	return &RemoteAlertmanagerMock_GetAlerts_Call{Call: _e.mock.On("GetAlerts", ctx, active, silenced, inhibited, filter, receiver)}
}

func (_c *RemoteAlertmanagerMock_GetAlerts_Call) Run(run func(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string)) *RemoteAlertmanagerMock_GetAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(bool), args[3].(bool), args[4].([]string), args[5].(string))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_GetAlerts_Call) Return(_a0 v2models.GettableAlerts, _a1 error) *RemoteAlertmanagerMock_GetAlerts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_GetAlerts_Call) RunAndReturn(run func(context.Context, bool, bool, bool, []string, string) (v2models.GettableAlerts, error)) *RemoteAlertmanagerMock_GetAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// GetReceivers provides a mock function with given fields: ctx
func (_m *RemoteAlertmanagerMock) GetReceivers(ctx context.Context) ([]v2models.Receiver, error) {
	ret := _m.Called(ctx)

	var r0 []v2models.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]v2models.Receiver, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []v2models.Receiver); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v2models.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_GetReceivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReceivers'
type RemoteAlertmanagerMock_GetReceivers_Call struct {
	*mock.Call
}

// GetReceivers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RemoteAlertmanagerMock_Expecter) GetReceivers(ctx interface{}) *RemoteAlertmanagerMock_GetReceivers_Call {
	return &RemoteAlertmanagerMock_GetReceivers_Call{Call: _e.mock.On("GetReceivers", ctx)}
}

func (_c *RemoteAlertmanagerMock_GetReceivers_Call) Run(run func(ctx context.Context)) *RemoteAlertmanagerMock_GetReceivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_GetReceivers_Call) Return(_a0 []v2models.Receiver, _a1 error) *RemoteAlertmanagerMock_GetReceivers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_GetReceivers_Call) RunAndReturn(run func(context.Context) ([]v2models.Receiver, error)) *RemoteAlertmanagerMock_GetReceivers_Call {
	_c.Call.Return(run)
	return _c
}

// GetSilence provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) GetSilence(_a0 context.Context, _a1 string) (v2models.GettableSilence, error) {
	ret := _m.Called(_a0, _a1)

	var r0 v2models.GettableSilence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (v2models.GettableSilence, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) v2models.GettableSilence); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(v2models.GettableSilence)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_GetSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSilence'
type RemoteAlertmanagerMock_GetSilence_Call struct {
	*mock.Call
}

// GetSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *RemoteAlertmanagerMock_Expecter) GetSilence(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_GetSilence_Call {
	return &RemoteAlertmanagerMock_GetSilence_Call{Call: _e.mock.On("GetSilence", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_GetSilence_Call) Run(run func(_a0 context.Context, _a1 string)) *RemoteAlertmanagerMock_GetSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_GetSilence_Call) Return(_a0 v2models.GettableSilence, _a1 error) *RemoteAlertmanagerMock_GetSilence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_GetSilence_Call) RunAndReturn(run func(context.Context, string) (v2models.GettableSilence, error)) *RemoteAlertmanagerMock_GetSilence_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields:
func (_m *RemoteAlertmanagerMock) GetStatus() definitions.GettableStatus {
	ret := _m.Called()

	var r0 definitions.GettableStatus
	if rf, ok := ret.Get(0).(func() definitions.GettableStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(definitions.GettableStatus)
	}

	return r0
}

// RemoteAlertmanagerMock_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type RemoteAlertmanagerMock_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
func (_e *RemoteAlertmanagerMock_Expecter) GetStatus() *RemoteAlertmanagerMock_GetStatus_Call {
	return &RemoteAlertmanagerMock_GetStatus_Call{Call: _e.mock.On("GetStatus")}
}

func (_c *RemoteAlertmanagerMock_GetStatus_Call) Run(run func()) *RemoteAlertmanagerMock_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_GetStatus_Call) Return(_a0 definitions.GettableStatus) *RemoteAlertmanagerMock_GetStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_GetStatus_Call) RunAndReturn(run func() definitions.GettableStatus) *RemoteAlertmanagerMock_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// ListSilences provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) ListSilences(_a0 context.Context, _a1 []string) (v2models.GettableSilences, error) {
	ret := _m.Called(_a0, _a1)

	var r0 v2models.GettableSilences
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (v2models.GettableSilences, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) v2models.GettableSilences); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.GettableSilences)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_ListSilences_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSilences'
type RemoteAlertmanagerMock_ListSilences_Call struct {
	*mock.Call
}

// ListSilences is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []string
func (_e *RemoteAlertmanagerMock_Expecter) ListSilences(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_ListSilences_Call {
	return &RemoteAlertmanagerMock_ListSilences_Call{Call: _e.mock.On("ListSilences", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_ListSilences_Call) Run(run func(_a0 context.Context, _a1 []string)) *RemoteAlertmanagerMock_ListSilences_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_ListSilences_Call) Return(_a0 v2models.GettableSilences, _a1 error) *RemoteAlertmanagerMock_ListSilences_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_ListSilences_Call) RunAndReturn(run func(context.Context, []string) (v2models.GettableSilences, error)) *RemoteAlertmanagerMock_ListSilences_Call {
	_c.Call.Return(run)
	return _c
}

// PutAlerts provides a mock function with given fields: _a0, _a1
func (_m *RemoteAlertmanagerMock) PutAlerts(_a0 context.Context, _a1 definitions.PostableAlerts) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.PostableAlerts) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_PutAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutAlerts'
type RemoteAlertmanagerMock_PutAlerts_Call struct {
	*mock.Call
}

// PutAlerts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 definitions.PostableAlerts
func (_e *RemoteAlertmanagerMock_Expecter) PutAlerts(_a0 interface{}, _a1 interface{}) *RemoteAlertmanagerMock_PutAlerts_Call {
	return &RemoteAlertmanagerMock_PutAlerts_Call{Call: _e.mock.On("PutAlerts", _a0, _a1)}
}

func (_c *RemoteAlertmanagerMock_PutAlerts_Call) Run(run func(_a0 context.Context, _a1 definitions.PostableAlerts)) *RemoteAlertmanagerMock_PutAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.PostableAlerts))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_PutAlerts_Call) Return(_a0 error) *RemoteAlertmanagerMock_PutAlerts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_PutAlerts_Call) RunAndReturn(run func(context.Context, definitions.PostableAlerts) error) *RemoteAlertmanagerMock_PutAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *RemoteAlertmanagerMock) Ready() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoteAlertmanagerMock_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type RemoteAlertmanagerMock_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *RemoteAlertmanagerMock_Expecter) Ready() *RemoteAlertmanagerMock_Ready_Call {
	return &RemoteAlertmanagerMock_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *RemoteAlertmanagerMock_Ready_Call) Run(run func()) *RemoteAlertmanagerMock_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_Ready_Call) Return(_a0 bool) *RemoteAlertmanagerMock_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_Ready_Call) RunAndReturn(run func() bool) *RemoteAlertmanagerMock_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAndApplyConfig provides a mock function with given fields: ctx, config
func (_m *RemoteAlertmanagerMock) SaveAndApplyConfig(ctx context.Context, config *definitions.PostableUserConfig) error {
	ret := _m.Called(ctx, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *definitions.PostableUserConfig) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_SaveAndApplyConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAndApplyConfig'
type RemoteAlertmanagerMock_SaveAndApplyConfig_Call struct {
	*mock.Call
}

// SaveAndApplyConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - config *definitions.PostableUserConfig
func (_e *RemoteAlertmanagerMock_Expecter) SaveAndApplyConfig(ctx interface{}, config interface{}) *RemoteAlertmanagerMock_SaveAndApplyConfig_Call {
	return &RemoteAlertmanagerMock_SaveAndApplyConfig_Call{Call: _e.mock.On("SaveAndApplyConfig", ctx, config)}
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyConfig_Call) Run(run func(ctx context.Context, config *definitions.PostableUserConfig)) *RemoteAlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*definitions.PostableUserConfig))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyConfig_Call) Return(_a0 error) *RemoteAlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyConfig_Call) RunAndReturn(run func(context.Context, *definitions.PostableUserConfig) error) *RemoteAlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAndApplyDefaultConfig provides a mock function with given fields: ctx
func (_m *RemoteAlertmanagerMock) SaveAndApplyDefaultConfig(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAndApplyDefaultConfig'
type RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call struct {
	*mock.Call
}

// SaveAndApplyDefaultConfig is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RemoteAlertmanagerMock_Expecter) SaveAndApplyDefaultConfig(ctx interface{}) *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	return &RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call{Call: _e.mock.On("SaveAndApplyDefaultConfig", ctx)}
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call) Run(run func(ctx context.Context)) *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call) Return(_a0 error) *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call) RunAndReturn(run func(context.Context) error) *RemoteAlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Return(run)
	return _c
}

// SendStateAndConfig provides a mock function with given fields:
func (_m *RemoteAlertmanagerMock) SendStateAndConfig() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoteAlertmanagerMock_SendStateAndConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendStateAndConfig'
type RemoteAlertmanagerMock_SendStateAndConfig_Call struct {
	*mock.Call
}

// SendStateAndConfig is a helper method to define mock.On call
func (_e *RemoteAlertmanagerMock_Expecter) SendStateAndConfig() *RemoteAlertmanagerMock_SendStateAndConfig_Call {
	return &RemoteAlertmanagerMock_SendStateAndConfig_Call{Call: _e.mock.On("SendStateAndConfig")}
}

func (_c *RemoteAlertmanagerMock_SendStateAndConfig_Call) Run(run func()) *RemoteAlertmanagerMock_SendStateAndConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_SendStateAndConfig_Call) Return(_a0 error) *RemoteAlertmanagerMock_SendStateAndConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoteAlertmanagerMock_SendStateAndConfig_Call) RunAndReturn(run func() error) *RemoteAlertmanagerMock_SendStateAndConfig_Call {
	_c.Call.Return(run)
	return _c
}

// StopAndWait provides a mock function with given fields:
func (_m *RemoteAlertmanagerMock) StopAndWait() {
	_m.Called()
}

// RemoteAlertmanagerMock_StopAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopAndWait'
type RemoteAlertmanagerMock_StopAndWait_Call struct {
	*mock.Call
}

// StopAndWait is a helper method to define mock.On call
func (_e *RemoteAlertmanagerMock_Expecter) StopAndWait() *RemoteAlertmanagerMock_StopAndWait_Call {
	return &RemoteAlertmanagerMock_StopAndWait_Call{Call: _e.mock.On("StopAndWait")}
}

func (_c *RemoteAlertmanagerMock_StopAndWait_Call) Run(run func()) *RemoteAlertmanagerMock_StopAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_StopAndWait_Call) Return() *RemoteAlertmanagerMock_StopAndWait_Call {
	_c.Call.Return()
	return _c
}

func (_c *RemoteAlertmanagerMock_StopAndWait_Call) RunAndReturn(run func()) *RemoteAlertmanagerMock_StopAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// TestReceivers provides a mock function with given fields: ctx, c
func (_m *RemoteAlertmanagerMock) TestReceivers(ctx context.Context, c definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error) {
	ret := _m.Called(ctx, c)

	var r0 *notifier.TestReceiversResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestReceiversConfigBodyParams) *notifier.TestReceiversResult); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notifier.TestReceiversResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, definitions.TestReceiversConfigBodyParams) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_TestReceivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestReceivers'
type RemoteAlertmanagerMock_TestReceivers_Call struct {
	*mock.Call
}

// TestReceivers is a helper method to define mock.On call
//   - ctx context.Context
//   - c definitions.TestReceiversConfigBodyParams
func (_e *RemoteAlertmanagerMock_Expecter) TestReceivers(ctx interface{}, c interface{}) *RemoteAlertmanagerMock_TestReceivers_Call {
	return &RemoteAlertmanagerMock_TestReceivers_Call{Call: _e.mock.On("TestReceivers", ctx, c)}
}

func (_c *RemoteAlertmanagerMock_TestReceivers_Call) Run(run func(ctx context.Context, c definitions.TestReceiversConfigBodyParams)) *RemoteAlertmanagerMock_TestReceivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.TestReceiversConfigBodyParams))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_TestReceivers_Call) Return(_a0 *notifier.TestReceiversResult, _a1 error) *RemoteAlertmanagerMock_TestReceivers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_TestReceivers_Call) RunAndReturn(run func(context.Context, definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error)) *RemoteAlertmanagerMock_TestReceivers_Call {
	_c.Call.Return(run)
	return _c
}

// TestTemplate provides a mock function with given fields: ctx, c
func (_m *RemoteAlertmanagerMock) TestTemplate(ctx context.Context, c definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error) {
	ret := _m.Called(ctx, c)

	var r0 *notify.TestTemplatesResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestTemplatesConfigBodyParams) *notify.TestTemplatesResults); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notify.TestTemplatesResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, definitions.TestTemplatesConfigBodyParams) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAlertmanagerMock_TestTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestTemplate'
type RemoteAlertmanagerMock_TestTemplate_Call struct {
	*mock.Call
}

// TestTemplate is a helper method to define mock.On call
//   - ctx context.Context
//   - c definitions.TestTemplatesConfigBodyParams
func (_e *RemoteAlertmanagerMock_Expecter) TestTemplate(ctx interface{}, c interface{}) *RemoteAlertmanagerMock_TestTemplate_Call {
	return &RemoteAlertmanagerMock_TestTemplate_Call{Call: _e.mock.On("TestTemplate", ctx, c)}
}

func (_c *RemoteAlertmanagerMock_TestTemplate_Call) Run(run func(ctx context.Context, c definitions.TestTemplatesConfigBodyParams)) *RemoteAlertmanagerMock_TestTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.TestTemplatesConfigBodyParams))
	})
	return _c
}

func (_c *RemoteAlertmanagerMock_TestTemplate_Call) Return(_a0 *notify.TestTemplatesResults, _a1 error) *RemoteAlertmanagerMock_TestTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoteAlertmanagerMock_TestTemplate_Call) RunAndReturn(run func(context.Context, definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error)) *RemoteAlertmanagerMock_TestTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// NewRemoteAlertmanagerMock creates a new instance of RemoteAlertmanagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRemoteAlertmanagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RemoteAlertmanagerMock {
	mock := &RemoteAlertmanagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
