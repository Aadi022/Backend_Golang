package service

import (
	"errors"

	"github.com/Aadi022/Backend_Golang/internal/model"
	"github.com/Aadi022/Backend_Golang/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository //repository instance that talks to the user db
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(name, email, password string) error {
	//First checks if a user with same email exists in the User table
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}

	//bcrypt works with byte slice. it converts string to byte slice to bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword( //GenerateFromPassword is the function that hashes a password
		[]byte(password),   //converts string to byte slice
		bcrypt.DefaultCost, //DefaultCost tells bcrypt how much work to do while hashing
	)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
}
