package users

import "log"

// capa de Servicio | logica de negocio
type (
	Service interface {
		Create(firstName, lastName, email, phone string) (*User, error)
		Get(id string) (*User, error)
		GetAll(filters Filters) ([]User, error)
		Delete(id string) error
		Update(id string, firstName, lastName, email, phone *string) (*User, error)
		Count(filters Filters) (int, error)
	}
	service struct {
		log  *log.Logger
		repo Respository
	}
	Filters struct {
		FirstName string
		LastName  string
	}
)

func NewService(log *log.Logger, repo Respository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}
func (s service) Create(firstName, lastName, email, phone string) (*User, error) {
	s.log.Println("Create user service")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	if err := s.repo.Create(&user); err != nil {
		s.log.Println(err)
		return nil, err
	}
	return &user, nil
}
func (s service) Get(id string) (*User, error) {
	users, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s service) GetAll(filters Filters) ([]User, error) {

	users, err := s.repo.GetAll(filters)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s service) Delete(id string) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
func (s service) Update(id string, firstName, lastName, email, phone *string) (*User, error) {
	err := s.repo.Update(id, firstName, lastName, email, phone)
	if err != nil {
		return nil, err
	}
	user, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
