package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
    ID       uint           `json:"id" gorm:"primaryKey"`
    Username     string         `json:"username"`
    Useremail    string         `json:"useremail" gorm:email"unique;not null"`
    Password string         `json:"password"`
    CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
