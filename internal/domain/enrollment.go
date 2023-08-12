package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Enrollment struct {
	ID        string     `xml:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string     `xml:"user_id,omitempty" gorm:"type:char(36)"`
	User      *User      `xml:"user,omitempty"`
	CourseID  string     `xml:"course_id" gorm:"type:char(36);not null"`
	Course    *Course    `xml:"course,omitempty"`
	Status    string     `xml:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time `xml:"-"`
	UpdatedAt *time.Time `xml:"-"`
}

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
