package enrollment

import (
	"errors"
	"gocourse/internal/courses"
	"gocourse/internal/domain"
	"gocourse/internal/users"
	"log"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}
	service struct {
		log       *log.Logger
		userSrv   users.Service
		courseSrv courses.Service
		repo      Repository
	}
)

func NewService(l *log.Logger, userSrv users.Service, courseSrv courses.Service, repo Repository) Service {
	return &service{
		log:       l,
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}
func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}
	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		return nil, errors.New("usuario no existe")
	}
	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil {
		return nil, errors.New("curso no existe")
	}
	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}
	return enroll, nil
}
