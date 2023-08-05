package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/main/internal/config"
	"backend/main/internal/models"
)

func CreateSchema(store models.DataStore, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema models.Schema
		if err := c.ShouldBindJSON(&schema); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := store.CreateSchema(context.Background(), schema); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, schema)
	}
}

func GetSchema(store models.DataStore, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		schema, err := store.GetSchema(context.Background(), name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, schema)
	}
}

func CreateRecord(store models.DataStore, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")

		var record map[string]interface{}
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := store.CreateRecord(context.Background(), collection, record); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, record)
	}
}

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
