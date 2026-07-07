package service

import (
	"errors"

	"github.com/Aadi022/Backend_Golang/internal/model"
	"github.com/Aadi022/Backend_Golang/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository //repository instance that talks to the user db
	jwt  *JWTService
}

func NewUserService(repo *repository.UserRepository, jwt *JWTService) *UserService {
	return &UserService{
		repo: repo,
		jwt:  jwt,
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

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email) //Find the user by email
	if err != nil {
		return "", errors.New("invalid email or password") //common practice to not reveal which field is invalid due to security reasons
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) //compares the entered password with the hashed password using bcrypt's function
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token, err := s.jwt.Generate(user.ID) //generate a jwt containing the user's ID
	if err != nil {
		return "", err
	}

	return token, nil //return the jwt token
}
