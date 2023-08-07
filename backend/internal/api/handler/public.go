package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/main/internal/auth"
	"backend/main/internal/config"
	"backend/main/internal/models"
)

func RegisterUser(auth auth.IAuthService, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = auth.RegisterUser(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user registered"})
	}
}

func LoginUser(auth auth.IAuthService, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := auth.AuthenticateUser(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

		// set jwt token in cookie
		c.SetCookie("jwt", token, (config.Authentication.TokenExpire * int(time.Minute)), "/", c.Request.Host, false, true)
	}
}

func LogoutUser(auth auth.IAuthService, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete jwt token from cookie
		c.SetCookie("jwt", "", -1, "/", c.Request.Host, false, true)

		c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
	}
}

func ValidateToken(ts auth.ITokenService, config *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        
        _, err := ts.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
    }
}

func RefreshToken(ts auth.ITokenService, config *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        refreshToken := c.GetHeader("Authorization")

		tokenDuration := time.Duration(config.Authentication.TokenExpire) * time.Minute
        newToken, err := ts.RefreshToken(refreshToken, tokenDuration)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"token": newToken})
    }
}
