package enrollment

import (
	"gocourse/internal/domain"
	"gorm.io/gorm"
	"log"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}
	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}
func (r *repo) Create(enroll *domain.Enrollment) error {
	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Println(err.Error())
		return err
	}
	r.log.Println("Enrollment created with id: ", enroll.ID)
	return nil
}
