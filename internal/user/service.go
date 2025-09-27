// File: internal/user/service.go

package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (*User, *ResponseError)
	LoginUser(input LoginUserInput) (string, *ResponseError)
	GetUserByID(id uint) (*User, *ResponseError)
}

type service struct {
	repo Repository;
	JWT_SECRET []byte;
}

func NewService(repo Repository, jwt_secret string) Service {
	return &service{repo: repo, JWT_SECRET: []byte(jwt_secret)}
}

func (s *service) RegisterUser(input RegisterUserInput) (*User, *ResponseError) {
	_, err := s.repo.FindByEmail(input.Email)
	if err == nil {
		return nil, ErrEmailExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternalServer
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternalServer
	}

	newUser := &User{
		Name:       input.Name,
		Email:      input.Email,
		Password:   string(hashedPassword),
		Identifier: input.Identifier,
		Address:    input.Address,
		Role:       input.Role,
	}

	if err := s.repo.Create(newUser); err != nil {
		return nil, ErrInternalServer
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (string, *ResponseError) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
		fmt.Println(err)
		return "", ErrInternalServer
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	claims := &Claims{
		UserID:     user.ID,
		Identifier: user.Identifier,
		Role:       user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.JWT_SECRET)
	if err != nil {
		fmt.Println(err)
		return "", ErrInternalServer
	}

	return signedToken, nil
}

func (s *service) GetUserByID(id uint) (*User, *ResponseError) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrInternalServer
	}
	return user, nil
}