package reply

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/moira-alert/moira-alert"
	"github.com/moira-alert/moira-alert/database"
)

func Notification(rep interface{}, err error) (moira.ScheduledNotification, error) {
	notification := moira.ScheduledNotification{}
	bytes, err := redis.Bytes(rep, err)
	if err != nil {
		return notification, err
	}
	err = json.Unmarshal(bytes, &notification)
	if err != nil {
		return notification, fmt.Errorf("Failed to parse notification json %s: %s", string(bytes), err.Error())
	}
	return notification, nil
}

func Notifications(rep interface{}, err error) ([]*moira.ScheduledNotification, error) {
	values, err := redis.Values(rep, err)
	if err != nil {
		if err == redis.ErrNil {
			return make([]*moira.ScheduledNotification, 0), nil
		}
		return nil, err
	}
	notifications := make([]*moira.ScheduledNotification, len(values))
	for i, value := range values {
		notification, err2 := Notification(value, err)
		if err2 != nil && err2 != database.ErrNil {
			return nil, err2
		} else if err2 == database.ErrNil {
			notifications[i] = nil
		} else {
			notifications[i] = &notification
		}
	}
	return notifications, nil
}