package user

import (
	"errors"
	"project-api/internal/modules/auth"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(data RegisterDTO) error {
	exists, _ := s.repo.FindByEmail(data.Email)

	if exists != nil {
		return errors.New("email already exists")
	}

	hash, err := auth.HashPassword(data.Password)

	if err != nil {
		return err
	}

	user := &User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
	}

	s.repo.Create(user)


	if err != nil {
		
		return err
	}

	return nil
}



func (s *Service) Login(data LoginDTO) (string, error) {
	user, _ := s.repo.FindByEmail(data.Email)

	if user == nil {
		return "", errors.New("invalid credentials")
	}

	valid := auth.CheckPassword(data.Password, user.Password)

	if !valid {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Delete(id uuid.UUID) error{

	user, err := s.repo.FindByID(id)

	if err != nil {
		return err
	}

	return s.repo.Delete(user.ID)
}

func (s *Service) FindAll() ([]User, error) {

	return s.repo.FindAll()
}