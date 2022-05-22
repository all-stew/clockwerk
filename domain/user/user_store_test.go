package user

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"fmt"
	"testing"
)

func TestStore_Create(t *testing.T) {

	if err := config.Init("./config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		username  string
		nickname  string
		email     string
		phone     string
		createdBy uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{"zjj", "zjj", "zjj", "zjj", 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Store{}
			if got := st.Create(tt.args.username, tt.args.nickname, tt.args.email, tt.args.phone, tt.args.createdBy); got != tt.want {
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
		id        uint64
		nickname  string
		email     string
		phone     string
		updatedBy uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test", args{103, "zjj", "zjj", "zjj", 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Store{}
			if got := st.Update(tt.args.id, tt.args.nickname, tt.args.email, tt.args.phone, tt.args.updatedBy); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
