package handler

import (
	"context"
	"net/http"

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

		token, err := auth.LoginUser(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
