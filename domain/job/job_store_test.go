package job

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"testing"
)

func TestStore_Create(t *testing.T) {

	if err := config.Init("../../myconfig.yaml"); err != nil {
		logger.Logf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		logger.Logf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		accountId          uint64
		userNotificationId uint64
		userId             uint64
		timeWindow         int
		createdBy          uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{accountId: 1, userNotificationId: 1, timeWindow: 1, createdBy: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.Create(tt.args.accountId, tt.args.userNotificationId, tt.args.userId, tt.args.timeWindow, tt.args.createdBy); got != tt.want {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
