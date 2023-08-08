package courses

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string         `xml:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name      string         `xml:"name" gorm:"type:char(50);not null"`
	StartDate time.Time      `xml:"start_date"`
	EndDate   time.Time      `xml:"end_date" `
	CreatedAt time.Time      `xml:"-"`
	UpdatedAt time.Time      `xml:"-"`
	Delete    gorm.DeletedAt `xml:"-"`
}
