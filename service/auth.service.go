package service

import (
	"errors"
	"log"

	"github.com/ydhnwb/golang_heroku/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) error
	// CreateUser(user dto.) entity.User
}

type authService struct {
	userRepo repo.UserRepository
}

func NewAuthService(userRepo repo.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (c *authService) VerifyCredential(email string, password string) error {
	user, err := c.userRepo.FindByEmail(email)
	if err != nil {
		println("hehe")
		println(err.Error())
		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
