package models

import "clockwerk/pkg/model"

type SysMenu struct {
	model.BaseModel
	MenuName   string      `json:"menu_name" gorm:"size:50;"`
	ParentId   uint64      `json:"parent_id" gorm:"not null;"`
	OrderNum   int         `json:"order_num" gorm:""`
	Path       string      `json:"path" gorm:"size:128;"`
	Component  string      `json:"component" gorm:"size:255;"`
	Query      string      `json:"query" gorm:"size:255;"`
	IsFrame    bool        `json:"is_frame" gorm:""`
	IsCache    bool        `json:"is_cache" gorm:""`
	MenuType   SysMenuType `json:"menuType" gorm:"size:1;"`
	IsVisible  bool        `json:"is_visible" gorm:""`
	Permission string      `json:"permission" gorm:"size:255;"`
	Status     int         `json:"status" gorm:""`
	Icon       string      `json:"icon" gorm:"size:128;"`
	Children   []SysMenu   `json:"children" gorm:"-"`
}

type SysMenuType int

const (
	SYS_MENU_TYPE_MENU     SysMenuType = 0
	SYS_MENU_TYPE_CATAGORY SysMenuType = 1
	SYS_MENU_TYPE_BUTTON   SysMenuType = 2
)

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (e *SysMenu) Generate() model.Builder {
	o := *e
	return &o
}

func (e *SysMenu) GetId() interface{} {
	return e.Id
}
