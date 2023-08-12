package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string         `xml:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name      string         `xml:"name" gorm:"type:char(50);not null"`
	StartDate time.Time      `xml:"start_date"`
	EndDate   time.Time      `xml:"end_date" `
	User      *User          `gorm:"-"`
	CreatedAt time.Time      `xml:"-"`
	UpdatedAt time.Time      `xml:"-"`
	Delete    gorm.DeletedAt `xml:"-"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
