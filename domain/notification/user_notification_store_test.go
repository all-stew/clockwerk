package notification

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"fmt"
	"testing"
)

func TestStore_Create(t *testing.T) {

	if err := config.Init("../../myconfig.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		notificationType  NotificationType
		notificationKey   string
		notificationParam string
		userId            uint64
		createdBy         uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{notificationParam: "{'test':'test'}", notificationKey: "123", notificationType: NOTIFICATION_TYPE_NONE, userId: 1, createdBy: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.Create(tt.args.notificationType, tt.args.notificationKey, tt.args.notificationParam, tt.args.userId, tt.args.createdBy); got != tt.want {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Update(t *testing.T) {

	if err := config.Init("../../myconfig.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		id                uint64
		notificationType  NotificationType
		notificationKey   string
		notificationParam string
		userId            uint64
		updatedBy         uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{id: 100, notificationParam: "{'test':'test1'}", notificationKey: "123", notificationType: NOTIFICATION_TYPE_NONE, userId: 1, updatedBy: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.Update(tt.args.id, tt.args.notificationType, tt.args.notificationKey, tt.args.notificationParam, tt.args.userId, tt.args.updatedBy); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
