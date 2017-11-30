package models

import (
	"time"
)

// Comment Model is the Data Structure for Database Migration and Querying
// It extends gorm.Model structure and allow us to execute operations
// on it through gORM instance.
type Comment struct {
	Id           uint64     `gorm:"primary_key" json:"id,omitempty"`
	Content      string     `gorm:"size:2000" json:"content"`
	Author       string     `gorm:"size:100" json:"author"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// TableName returns and indicate to gORM the model table name
// It allows us and gORM to retrieve the exact table name
func (Comment) TableName() string {
	return "comments"
}
