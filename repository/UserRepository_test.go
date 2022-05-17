package repository

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {

	if err := config.Init("../myconfig.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		username string
		nickname string
		email    string
		phone    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"zjj", "zjj", "zjj", "zjj"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.username, tt.args.nickname, tt.args.email, tt.args.phone); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	if err := config.Init("../myconfig.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	type args struct {
		id       uint64
		username string
		nickname string
		email    string
		phone    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{100, "zjj1", "zjj1", "zjj", "zjj"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Update(tt.args.id, tt.args.username, tt.args.nickname, tt.args.email, tt.args.phone); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
