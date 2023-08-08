package courses

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(course *Course) error
		Get(id string) (*Course, error)
		GetAll(filters Filters) ([]Course, error)
		Delete(id string) error
		Update(id string, name, startDate, endDate *string) error
		Count(filters Filters) (int, error)
	}
	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func (r *repo) Create(course *Course) error {
	if err := r.db.Create(course).Error; err != nil {
		r.log.Println(err)
		return err
	}
	r.log.Println("repository and user with ", course.ID)
	return nil
}
func NewRepo(log *log.Logger, db *gorm.DB) (Repository, error) {
	return &repo{
		log: log,
		db:  db,
	}, nil
}
func (r *repo) Get(id string) (*Course, error) {
	course := Course{ID: id}
	if err := r.db.First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
func (r *repo) GetAll(filters Filters) ([]Course, error) {
	var c []Course
	tx := r.db.Model(&c)
	tx = applyFilters(tx, filters)
	if err := tx.Order("created_at desc").Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil

}
func (r *repo) Delete(id string) error {
	course := Course{ID: id}
	if err := r.db.Delete(&course).Error; err != nil {
		return err
	}
	return nil
}
func (r *repo) Update(id string, name, startDate, endDate *string) error {
	course := Course{ID: id}
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}
	if startDate != nil {
		values["startDate"] = *startDate
	}
	if endDate != nil {
		values["endDate"] = *endDate
	}

	if err := r.db.Model(&course).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	return tx
}
