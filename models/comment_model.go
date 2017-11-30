package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Comment Model is the Data Structure for Database Migration and Querying
// It extends gorm.Model structure and allow us to execute operations
// on it through gORM instance.
type Comment struct {
	gorm.Model
	Id           uint64      `gorm:"primary_key"`
	Content      string     `gorm:"size:2000"`
	Author       string     `gorm:"size:100"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// TableName returns and indicate to gORM the model table name
// It allows us and gORM to retrieve the exact table name
func (Comment) TableName() string {
	return "comments"
}
