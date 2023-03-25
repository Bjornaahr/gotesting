package middleware

import (
	"VekterBackend/src/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		token, err := controllers.ParseToken(tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, isValid := token.Claims.(jwt.MapClaims)

		if isValid && token.Valid {
			context.Set("email", claims["email"])
		} else {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
