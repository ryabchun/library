package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(name, email, password, role string) (*User, error)
	LoginUser(email, password string) (*User, error)
	GetUser(id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) RegisterUser(name, email, password, role string) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Name:     name,
		Email:    email,
		Password: string(hashed), // хэш
		Role:     role,
	}

	if err := s.repo.CreateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) LoginUser(email, password string) (*User, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}

func (s *service) GetUser(id uint) (*User, error) {
	return s.repo.GetByID(id)
}
