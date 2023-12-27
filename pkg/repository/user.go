package repository

import (
	"GoBookShelf/pkg/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user models.User) (*models.User, error)
	CheckEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context,email string)(models.User,error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(ctx context.Context, user models.User) (*models.User, error) {
	// err := r.db.Create(&user).Error
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "user_repository",
			"user":   user.Email,
			"error":  err,
			"where":  "RegisterUser(repository)",
		}).Error("Failed to Create User")

		logrus.Error("Error occurred during creating new user: ", err, ", User: ", user.Email)

		// log.Errorf("Error occured during creating new user: %v, Kullanıcı: %v", err, user.Email)
		return nil, errors.New("uunable to create user")

	}
	return &user, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, email string) (bool, error) {
	var user models.User
	if err := r.db.Where("email=?", email).Take(&user).Error; err != nil {
		return false, err

	}
	return true, nil

}


func(r*userRepository)GetUserByEmail(ctx context.Context,email string )(models.User ,error){
	var user models.User
	if err:=r.db.Where("email=?",email).Take(&user).Error; err!=nil{
		return models.User{},err
	}
	return user,nil
}
