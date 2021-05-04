package repo

import (
	"log"

	"github.com/ydhnwb/golang_heroku/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUserID(userID string) (entity.User, error)
}

type userRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection *gorm.DB) UserRepository {
	return &userRepo{
		connection: connection,
	}
}

func (c *userRepo) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	c.connection.Save(&user)
	return user, nil
}

func (c *userRepo) UpdateUser(user entity.User) (entity.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		c.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	c.connection.Save(&user)
	return user, nil
}

func (c *userRepo) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	res := c.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *userRepo) FindByUserID(userID string) (entity.User, error) {
	var user entity.User
	res := c.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
