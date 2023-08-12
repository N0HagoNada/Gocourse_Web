package courses

import (
	"gocourse/internal/domain"
	"log"
	"time"
)

const (
	DateOnly = "2006-01-02"
)

type (
	Service interface {
		Create(name, startDate, endDate string) (*domain.Course, error)
		Get(id string) (*domain.Course, error)
		GetAll(filters Filters) ([]domain.Course, error)
		Delete(id string) error
		Update(id string, name, startDate, endDate *string) (*domain.Course, error)
		Count(filters Filters) (int, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
	Filters struct {
		Name string
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}
func (s service) Create(name, startDate, endDate string) (*domain.Course, error) {
	startDateParsed, err := time.Parse(DateOnly, startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	endDateParsed, err := time.Parse(DateOnly, endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	s.log.Println("Create course service")
	course := domain.Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}
	if err := s.repo.Create(&course); err != nil {
		s.log.Println(err)
		return nil, err
	}
	return &course, nil
}
func (s service) Get(id string) (*domain.Course, error) {
	courses, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (s service) GetAll(filters Filters) ([]domain.Course, error) {

	courses, err := s.repo.GetAll(filters)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (s service) Delete(id string) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
func (s service) Update(id string, name, startDate, endDate *string) (*domain.Course, error) {
	err := s.repo.Update(id, name, startDate, endDate)
	if err != nil {
		return nil, err
	}
	course, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}
func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
