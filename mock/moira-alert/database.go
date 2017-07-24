// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira-alert (interfaces: Database)

package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	moira_alert "github.com/moira-alert/moira-alert"
	time "time"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return _m.recorder
}

// AddNotification mocks base method
func (_m *MockDatabase) AddNotification(_param0 *moira_alert.ScheduledNotification) error {
	ret := _m.ctrl.Call(_m, "AddNotification", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotification indicates an expected call of AddNotification
func (_mr *MockDatabaseMockRecorder) AddNotification(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddNotification", arg0)
}

// CreateSubscription mocks base method
func (_m *MockDatabase) CreateSubscription(_param0 *moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "CreateSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscription indicates an expected call of CreateSubscription
func (_mr *MockDatabaseMockRecorder) CreateSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateSubscription", arg0)
}

// DeleteContact mocks base method
func (_m *MockDatabase) DeleteContact(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteContact", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContact indicates an expected call of DeleteContact
func (_mr *MockDatabaseMockRecorder) DeleteContact(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteContact", arg0, arg1)
}

// DeleteSubscription mocks base method
func (_m *MockDatabase) DeleteSubscription(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSubscription", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (_mr *MockDatabaseMockRecorder) DeleteSubscription(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSubscription", arg0, arg1)
}

// DeleteTag mocks base method
func (_m *MockDatabase) DeleteTag(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTag", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag
func (_mr *MockDatabaseMockRecorder) DeleteTag(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTag", arg0)
}

// FetchEvent mocks base method
func (_m *MockDatabase) FetchEvent() (*moira_alert.EventData, error) {
	ret := _m.ctrl.Call(_m, "FetchEvent")
	ret0, _ := ret[0].(*moira_alert.EventData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEvent indicates an expected call of FetchEvent
func (_mr *MockDatabaseMockRecorder) FetchEvent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchEvent")
}

// GetAllContacts mocks base method
func (_m *MockDatabase) GetAllContacts() ([]moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetAllContacts")
	ret0, _ := ret[0].([]moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllContacts indicates an expected call of GetAllContacts
func (_mr *MockDatabaseMockRecorder) GetAllContacts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllContacts")
}

// GetChecksCount mocks base method
func (_m *MockDatabase) GetChecksCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetChecksCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecksCount indicates an expected call of GetChecksCount
func (_mr *MockDatabaseMockRecorder) GetChecksCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetChecksCount")
}

// GetContact mocks base method
func (_m *MockDatabase) GetContact(_param0 string) (moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetContact", _param0)
	ret0, _ := ret[0].(moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact
func (_mr *MockDatabaseMockRecorder) GetContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContact", arg0)
}

// GetEvents mocks base method
func (_m *MockDatabase) GetEvents(_param0 string, _param1 int64, _param2 int64) ([]*moira_alert.EventData, error) {
	ret := _m.ctrl.Call(_m, "GetEvents", _param0, _param1, _param2)
	ret0, _ := ret[0].([]*moira_alert.EventData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents
func (_mr *MockDatabaseMockRecorder) GetEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEvents", arg0, arg1, arg2)
}

// GetFilteredTriggersIds mocks base method
func (_m *MockDatabase) GetFilteredTriggersIds(_param0 []string, _param1 bool) ([]string, int64, error) {
	ret := _m.ctrl.Call(_m, "GetFilteredTriggersIds", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFilteredTriggersIds indicates an expected call of GetFilteredTriggersIds
func (_mr *MockDatabaseMockRecorder) GetFilteredTriggersIds(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFilteredTriggersIds", arg0, arg1)
}

// GetMetricsCount mocks base method
func (_m *MockDatabase) GetMetricsCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetMetricsCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricsCount indicates an expected call of GetMetricsCount
func (_mr *MockDatabaseMockRecorder) GetMetricsCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricsCount")
}

// GetNotificationTrigger mocks base method
func (_m *MockDatabase) GetNotificationTrigger(_param0 string) (moira_alert.TriggerData, error) {
	ret := _m.ctrl.Call(_m, "GetNotificationTrigger", _param0)
	ret0, _ := ret[0].(moira_alert.TriggerData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationTrigger indicates an expected call of GetNotificationTrigger
func (_mr *MockDatabaseMockRecorder) GetNotificationTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotificationTrigger", arg0)
}

// GetNotifications mocks base method
func (_m *MockDatabase) GetNotifications(_param0 int64) ([]*moira_alert.ScheduledNotification, error) {
	ret := _m.ctrl.Call(_m, "GetNotifications", _param0)
	ret0, _ := ret[0].([]*moira_alert.ScheduledNotification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications
func (_mr *MockDatabaseMockRecorder) GetNotifications(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotifications", arg0)
}

// GetPatterns mocks base method
func (_m *MockDatabase) GetPatterns() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatterns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatterns indicates an expected call of GetPatterns
func (_mr *MockDatabaseMockRecorder) GetPatterns() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatterns")
}

// GetSubscription mocks base method
func (_m *MockDatabase) GetSubscription(_param0 string) (moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscription", _param0)
	ret0, _ := ret[0].(moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription
func (_mr *MockDatabaseMockRecorder) GetSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscription", arg0)
}

// GetSubscriptions mocks base method
func (_m *MockDatabase) GetSubscriptions(_param0 []string) ([]moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscriptions", _param0)
	ret0, _ := ret[0].([]moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (_mr *MockDatabaseMockRecorder) GetSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscriptions", arg0)
}

// GetTag mocks base method
func (_m *MockDatabase) GetTag(_param0 string) (moira_alert.TagData, error) {
	ret := _m.ctrl.Call(_m, "GetTag", _param0)
	ret0, _ := ret[0].(moira_alert.TagData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag
func (_mr *MockDatabaseMockRecorder) GetTag(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTag", arg0)
}

// GetTagNames mocks base method
func (_m *MockDatabase) GetTagNames() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagNames indicates an expected call of GetTagNames
func (_mr *MockDatabaseMockRecorder) GetTagNames() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagNames")
}

// GetTagTriggerIds mocks base method
func (_m *MockDatabase) GetTagTriggerIds(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagTriggerIds", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagTriggerIds indicates an expected call of GetTagTriggerIds
func (_mr *MockDatabaseMockRecorder) GetTagTriggerIds(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagTriggerIds", arg0)
}

// GetTags mocks base method
func (_m *MockDatabase) GetTags(_param0 []string) (map[string]moira_alert.TagData, error) {
	ret := _m.ctrl.Call(_m, "GetTags", _param0)
	ret0, _ := ret[0].(map[string]moira_alert.TagData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (_mr *MockDatabaseMockRecorder) GetTags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTags", arg0)
}

// GetTagsSubscriptions mocks base method
func (_m *MockDatabase) GetTagsSubscriptions(_param0 []string) ([]moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetTagsSubscriptions", _param0)
	ret0, _ := ret[0].([]moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsSubscriptions indicates an expected call of GetTagsSubscriptions
func (_mr *MockDatabaseMockRecorder) GetTagsSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagsSubscriptions", arg0)
}

// GetTrigger mocks base method
func (_m *MockDatabase) GetTrigger(_param0 string) (*moira_alert.Trigger, error) {
	ret := _m.ctrl.Call(_m, "GetTrigger", _param0)
	ret0, _ := ret[0].(*moira_alert.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger
func (_mr *MockDatabaseMockRecorder) GetTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTrigger", arg0)
}

// GetTriggerEventsCount mocks base method
func (_m *MockDatabase) GetTriggerEventsCount(_param0 string, _param1 int64) int64 {
	ret := _m.ctrl.Call(_m, "GetTriggerEventsCount", _param0, _param1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTriggerEventsCount indicates an expected call of GetTriggerEventsCount
func (_mr *MockDatabaseMockRecorder) GetTriggerEventsCount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerEventsCount", arg0, arg1)
}

// GetTriggerIds mocks base method
func (_m *MockDatabase) GetTriggerIds() ([]string, int64, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTriggerIds indicates an expected call of GetTriggerIds
func (_mr *MockDatabaseMockRecorder) GetTriggerIds() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerIds")
}

// GetTriggerLastCheck mocks base method
func (_m *MockDatabase) GetTriggerLastCheck(_param0 string) (*moira_alert.CheckData, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerLastCheck", _param0)
	ret0, _ := ret[0].(*moira_alert.CheckData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerLastCheck indicates an expected call of GetTriggerLastCheck
func (_mr *MockDatabaseMockRecorder) GetTriggerLastCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerLastCheck", arg0)
}

// GetTriggerTags mocks base method
func (_m *MockDatabase) GetTriggerTags(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerTags", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerTags indicates an expected call of GetTriggerTags
func (_mr *MockDatabaseMockRecorder) GetTriggerTags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerTags", arg0)
}

// GetTriggerThrottlingTimestamps mocks base method
func (_m *MockDatabase) GetTriggerThrottlingTimestamps(_param0 string) (time.Time, time.Time) {
	ret := _m.ctrl.Call(_m, "GetTriggerThrottlingTimestamps", _param0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// GetTriggerThrottlingTimestamps indicates an expected call of GetTriggerThrottlingTimestamps
func (_mr *MockDatabaseMockRecorder) GetTriggerThrottlingTimestamps(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerThrottlingTimestamps", arg0)
}

// GetTriggersChecks mocks base method
func (_m *MockDatabase) GetTriggersChecks(_param0 []string) ([]moira_alert.TriggerChecks, error) {
	ret := _m.ctrl.Call(_m, "GetTriggersChecks", _param0)
	ret0, _ := ret[0].([]moira_alert.TriggerChecks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggersChecks indicates an expected call of GetTriggersChecks
func (_mr *MockDatabaseMockRecorder) GetTriggersChecks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggersChecks", arg0)
}

// GetUserContacts mocks base method
func (_m *MockDatabase) GetUserContacts(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserContacts", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserContacts indicates an expected call of GetUserContacts
func (_mr *MockDatabaseMockRecorder) GetUserContacts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserContacts", arg0)
}

// GetUserSubscriptionIds mocks base method
func (_m *MockDatabase) GetUserSubscriptionIds(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserSubscriptionIds", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscriptionIds indicates an expected call of GetUserSubscriptionIds
func (_mr *MockDatabaseMockRecorder) GetUserSubscriptionIds(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserSubscriptionIds", arg0)
}

// PushEvent mocks base method
func (_m *MockDatabase) PushEvent(_param0 *moira_alert.EventData, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "PushEvent", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushEvent indicates an expected call of PushEvent
func (_mr *MockDatabaseMockRecorder) PushEvent(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PushEvent", arg0, arg1)
}

// SaveMetrics mocks base method
func (_m *MockDatabase) SaveMetrics(_param0 map[string]*moira_alert.MatchedMetric) error {
	ret := _m.ctrl.Call(_m, "SaveMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMetrics indicates an expected call of SaveMetrics
func (_mr *MockDatabaseMockRecorder) SaveMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveMetrics", arg0)
}

// SetContact mocks base method
func (_m *MockDatabase) SetContact(_param0 *moira_alert.ContactData) error {
	ret := _m.ctrl.Call(_m, "SetContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContact indicates an expected call of SetContact
func (_mr *MockDatabaseMockRecorder) SetContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetContact", arg0)
}

// SetTagMaintenance mocks base method
func (_m *MockDatabase) SetTagMaintenance(_param0 string, _param1 moira_alert.TagData) error {
	ret := _m.ctrl.Call(_m, "SetTagMaintenance", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTagMaintenance indicates an expected call of SetTagMaintenance
func (_mr *MockDatabaseMockRecorder) SetTagMaintenance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTagMaintenance", arg0, arg1)
}

// SetTriggerMetricsMaintenance mocks base method
func (_m *MockDatabase) SetTriggerMetricsMaintenance(_param0 string, _param1 map[string]int64) error {
	ret := _m.ctrl.Call(_m, "SetTriggerMetricsMaintenance", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerMetricsMaintenance indicates an expected call of SetTriggerMetricsMaintenance
func (_mr *MockDatabaseMockRecorder) SetTriggerMetricsMaintenance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerMetricsMaintenance", arg0, arg1)
}

// SetTriggerThrottlingTimestamp mocks base method
func (_m *MockDatabase) SetTriggerThrottlingTimestamp(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetTriggerThrottlingTimestamp", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerThrottlingTimestamp indicates an expected call of SetTriggerThrottlingTimestamp
func (_mr *MockDatabaseMockRecorder) SetTriggerThrottlingTimestamp(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerThrottlingTimestamp", arg0, arg1)
}

// UpdateMetricsHeartbeat mocks base method
func (_m *MockDatabase) UpdateMetricsHeartbeat() error {
	ret := _m.ctrl.Call(_m, "UpdateMetricsHeartbeat")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetricsHeartbeat indicates an expected call of UpdateMetricsHeartbeat
func (_mr *MockDatabaseMockRecorder) UpdateMetricsHeartbeat() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMetricsHeartbeat")
}

// UpdateSubscription mocks base method
func (_m *MockDatabase) UpdateSubscription(_param0 *moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "UpdateSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscription indicates an expected call of UpdateSubscription
func (_mr *MockDatabaseMockRecorder) UpdateSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSubscription", arg0)
}

// WriteContact mocks base method
func (_m *MockDatabase) WriteContact(_param0 *moira_alert.ContactData) error {
	ret := _m.ctrl.Call(_m, "WriteContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteContact indicates an expected call of WriteContact
func (_mr *MockDatabaseMockRecorder) WriteContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteContact", arg0)
}

// WriteSubscriptions mocks base method
func (_m *MockDatabase) WriteSubscriptions(_param0 []*moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "WriteSubscriptions", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteSubscriptions indicates an expected call of WriteSubscriptions
func (_mr *MockDatabaseMockRecorder) WriteSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteSubscriptions", arg0)
}
