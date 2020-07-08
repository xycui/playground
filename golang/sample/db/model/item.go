package model

import "github.com/jinzhu/gorm"

type DataItem struct {
	gorm.Model
	Data string `gorm:"column:data;index:idx_data;type:varchar(255)"`
}

func (*DataItem) TableName() string {
	return "dataitems"
}
