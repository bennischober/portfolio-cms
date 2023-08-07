package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/main/internal/config"
	"backend/main/internal/models"
)

func GetSingleRecord(store models.DataStore, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")
		id := c.Param("id")

		record, err := store.GetSingleRecord(context.Background(), collection, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, record)
	}
}

func GetRecords(store models.DataStore, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")

		records, err := store.GetRecords(context.Background(), collection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, records)
	}
}
