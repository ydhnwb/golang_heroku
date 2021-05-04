package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/common/response"
	"github.com/ydhnwb/golang_heroku/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := response.BuildErrorResponse("Failed to process request", "No token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token := jwtService.ValidateToken(authHeader, c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			response := response.BuildErrorResponse("Error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
