package user

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twitter-clone/models"
	users "github.com/twitter-clone/user-management"
	"golang.org/x/crypto/bcrypt"
)

type profile struct {
	models.Profile
}

// new user

func NewProfile(p models.Profile) users.Profile {
	return &profile{p}
}

func (p *profile) Authenticate(ctx context.Context, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
func (p *profile) GenerateUserToken(ctx context.Context) (string, error) {

	ExpiresAt := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": p.Username,
		"email":    p.Email,
		"exp":      float64(ExpiresAt),
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		return "", error
	}

	return tokenString, nil

}

func (p *profile) Encrypt(ctx context.Context, password string) (*models.Profile, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	p.Password = string(pw)
	return &p.Profile, nil

}
func (p *profile) Validate(ctx context.Context) error {

	if p.Username == "" {
		return errors.New("Please provide username")
	}

	if p.Password == "" {
		return errors.New("password field can't be empty")
	}
	return nil

}
