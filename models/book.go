package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
    ID     uint           `json:"id" gorm:"primaryKey"`
    Title  string         `json:"title" gorm:"not null"`
    Author string         `json:"author" gorm:"not null"`
    Descripition string `json:"description" gorm:"type:text"` // Descrição do livro

    // Relacionamento com User
    UserID uint           `json:"user_id"`             // Chave estrangeira
    User   Users          `json:"-" gorm:"foreignKey:UserID"`

    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
