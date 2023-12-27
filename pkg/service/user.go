package service

import (
	"GoBookShelf/dto"
	"GoBookShelf/helpers"
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/repository"
	"context"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	Verify(ctx context.Context, email string, password string) (bool, error)
	// Diğer metodlar (GetUserByID, UpdateUser, DeleteUser vb.)

}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	// Burada userRepo üzerinden CreateUser metodunu çağırabilir ve ek iş mantığı uygulayabilirsiniz
	email, _ := s.userRepo.CheckEmail(ctx, req.Email)
	if email {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		IsAdmin:      false,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":   "RegusterUser(Service)",
			"error":     err.Error(),
			"email":     user.Email,
			"name":      user.Name,
			"timestamp": time.Now().Format(time.RFC3339),
		}).Error("Failed to register user(Service)")

		return dto.UserResponse{}, dto.ErrCreateUser
	}
	//burası doldurulmalıdır sonra
	// draftEmail, err := makeVerificationEmail(userReg.Email)
	// if err != nil {
	// 	return dto.UserResponse{}, err
	// }

	// err = utils.SendMail(userReg.Email, draftEmail["subject"], draftEmail["body"])
	// if err != nil {
	// 	return dto.UserResponse{}, err
	// }
	return dto.UserResponse{
		Name:       userReg.Name,
		Email:      userReg.Email,
		IsAdmin:    userReg.IsAdmin,
		IsVerified: false,
	}, nil

}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	emails, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByEmail
	}
	// 	UserResponse struct {
	// 	ID         string `json:"id"`
	// 	Name       string `json:"name"`
	// 	Email      string `json:"email"`
	// 	IsAdmin    bool   `json:"is_admin"`
	// 	IsVerified bool   `json:"is_verified"` // Eğer kullanıcı doğrulama işlemi varsa
	// }

	return dto.UserResponse{
		ID:         strconv.Itoa(emails.ID),
		Name:       emails.Name,
		Email:      emails.Email,
		IsAdmin:    emails.IsAdmin,
		IsVerified: false,
	}, nil
}

func (s *userService) Verify(ctx context.Context, email string, password string) (bool, error) {
	res, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return false, dto.ErrUserNotFound
	}
	if !res.IsVerified {
		return false, dto.ErrAccountNotVerified
	}
	checkPassword, err := helpers.CheckPassword(res.PasswordHash, []byte(password))
	if err != nil {
		return false, dto.ErrPasswordNotMatch
	}

	if res.Email == email && checkPassword {
		return true, nil
	}

	return false, dto.ErrEmailOrPassword

}
