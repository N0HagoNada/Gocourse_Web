package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// capa de base de datos
type User struct {
	ID        string    `xml:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	FirstName string    `xml:"name" gorm:"type:char(50);not null"`
	LastName  string    `xml:"lastName" gorm:"type:char(50);not null"`
	Email     string    `xml:"email" gorm:"type:char(50);not null"`
	Phone     string    `xml:"phone" gorm:"type:char(30)"`
	Course    *Course   `gorm:"-"`
	CreatedAt time.Time `xml:"-"`
	UpdatedAt time.Time `xml:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
