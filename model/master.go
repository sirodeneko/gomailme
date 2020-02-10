package model

import (
	"github.com/jinzhu/gorm"
)

// Master 发件人模型
type Master struct {
	gorm.Model
	User string `gorm:"size:1000"`
	Pass string
}


