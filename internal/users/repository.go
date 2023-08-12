package users

import (
	"fmt"
	"gocourse/internal/domain"
	"gorm.io/gorm"
	"log"
	"strings"
)

// capa de Persistencia | conexion a la base de datos
type (
	Respository interface {
		Create(user *domain.User) error
		GetAll(filters Filters) ([]domain.User, error)
		Get(id string) (*domain.User, error)
		Delete(id string) error
		Update(id string, firstName, lastName, email, phone *string) error
		Count(filters Filters) (int, error)
	}
	repo struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepo(log *log.Logger, db *gorm.DB) (Respository, error) {
	return &repo{
		log: log,
		db:  db,
	}, nil
}
func (r *repo) Create(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		r.log.Println(err)
		return err
	}
	r.log.Println("repository and user with ", user.ID)
	return nil
}
func (r *repo) GetAll(filters Filters) ([]domain.User, error) {
	var u []domain.User
	tx := r.db.Model(&u)
	tx = applyFilters(tx, filters)
	if err := tx.Order("created_at desc").Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil

}
func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}
	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}
	return tx
}
func (r *repo) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}
	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repo) Delete(id string) error {
	user := domain.User{ID: id}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
func (r *repo) Update(id string, firstName, lastName, email, phone *string) error {
	user := domain.User{ID: id}
	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}
	if lastName != nil {
		values["last_name"] = *lastName
	}
	if email != nil {
		values["email"] = *email
	}
	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&user).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(domain.User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
