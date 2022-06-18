package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/repo"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string, ctx *gin.Context) *jwt.Token
	ExtractUser(token string, userRepo repo.UserRepository) interface{}
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "system"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string, ctx *gin.Context) *jwt.Token {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t

}

func (j *jwtService) ExtractUser(token string, userRepo repo.UserRepository) interface{} {
	t := j.ValidateToken(token, &gin.Context{})

	if t == nil {
		return nil
	}

	fUserID := t.Claims.(jwt.MapClaims)["user_id"].(float64)
	userID := strconv.FormatFloat(fUserID, 'E', -1, 64)
	user, _ := userRepo.FindByUserID(userID)
	return user
}
