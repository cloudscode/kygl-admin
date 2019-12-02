package model

import (
	"time"
)

type CodelibraryItem struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	ParentId       int       `json:"parent_id"`
	Sortvalue      int       `json:"sortvalue"`
	CreateTime     time.Time `gorm:"type:time;column:create_time;not null;default:CURRENT_TIMESTAMP" json:"created_time,omitempty" example:"2019-07-10 0:39"`
	LastUpdateTime time.Time `gorm:"type:time;column:last_update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE" json:"last_update_time,omitempty" example:"2019-07-10 0:39"`
	Deleted        int       `json:"deleted"`
}

func (CodelibraryItem) TableName() string {
	return "tab_codelibrary_item"
}
