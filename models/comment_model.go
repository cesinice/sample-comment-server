package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Id           int64      `gorm:"primary_key"`
	Content      string     `gorm:"size:2000"`
	Author       string     `gorm:"size:100"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (Comment) TableName() string {
	return "comments"
}
