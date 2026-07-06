package repository

import (
	"github.com/Aadi022/Backend_Golang/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB //db connection object
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) error { //we are using (user *model.User) because model.User is the type(class), and user is the object. So in go for any db changes, we accept object of the model
	return r.db.Create(&user).Error //r.db is the database connection
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User //GORM will fill this struct by the fetch from DB

	err := r.db.Where("email= ?", email).First(&user).Error //"email=?" is the sql condition (? is the placeholder), email replaces ? safely and prevents sql injection
	//First(&user) fetches the first matching row, we use & so we give original copy to gorm, if any modifications required gorm can do it
	//The fetched row is put in user

	if err != nil {
		return nil, err
	}

	return &user, nil
}
