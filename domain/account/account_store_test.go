package account

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"testing"
)

func TestStore_Create(t *testing.T) {
	if err := config.Init("../../config.yaml"); err != nil {
		logger.Logf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		logger.Logf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		account      string
		accountType  AccountType
		parentId     uint64
		accountParam string
		userId       uint64
		createBy     uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{account: "123", accountType: ACCOUNT_TYPE_MIHOYO_BBS, parentId: 0, accountParam: "{\"cookie_token\":\"\", \"account_id\": \"\"}", userId: 1, createBy: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.Create(tt.args.account, tt.args.accountType, tt.args.parentId, tt.args.accountParam, tt.args.userId, tt.args.createBy); got != tt.want {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Update(t *testing.T) {

	if err := config.Init("../../config.yaml"); err != nil {
		logger.Logf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		logger.Logf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		id           uint64
		accountParam string
		updatedBy    uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{id: 1, accountParam: "{\"test\":1}", updatedBy: 0}, false},
		{"test1", args{id: 100, accountParam: "{\"test\":1}", updatedBy: 0}, true},
		{"test1", args{id: 100, accountParam: "test", updatedBy: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.Update(tt.args.id, tt.args.accountParam, tt.args.updatedBy); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_ListByStatus(t *testing.T) {

	if err := config.Init("../../config.yaml"); err != nil {
		logger.Logf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		logger.Logf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		status Status
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{status: ACCOUNT_ENABLE}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Store{}
			if got := s.ListByStatus(tt.args.status); len(got) != tt.want {
				t.Errorf("ListByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
