package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	Authenticate() gin.HandlerFunc
}

type AuthMiddleware struct {
	TokenService ITokenService
}

func NewAuthMiddleware(tokenService ITokenService) *AuthMiddleware {
	return &AuthMiddleware{TokenService: tokenService}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized, no token provided"})
			c.Abort()
			return
		}

		claims, err := m.TokenService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
