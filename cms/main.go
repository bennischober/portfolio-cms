package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

type Schema struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name     string      `json:"name"`
	DataType string      `json:"data_type"`
	Data     interface{} `json:"data"`
}

func main() {
	// MongoDB client
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		mongodbURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Router
	r := gin.Default()

	// Handlers
	r.POST("/api/schema", func(c *gin.Context) {
		createSchema(c, client)
	})
	r.GET("/api/schema/:name", func(c *gin.Context) {
		getSchema(c, client)
	})

	// Run server
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func createSchema(c *gin.Context, client *mongo.Client) {
	var schema Schema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// iterate over each field, if it's a file field, handle the file upload
	for i, field := range schema.Fields {
		if field.DataType == "file" {
			file, err := c.FormFile(field.Name)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Save the file
			path := filepath.Join("./files", file.Filename)
			if err := c.SaveUploadedFile(file, path); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// replace the file data in the schema with the path
			schema.Fields[i].Data = path
		}
	}

	// Now save the schema to MongoDB
	collection := client.Database("test_db").Collection("test_col")
	if _, err := collection.InsertOne(context.Background(), schema); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func getSchema(c *gin.Context, client *mongo.Client) {
	name := c.Param("name")
	collection := client.Database("test_db").Collection("test_col")

	var schema Schema
	if err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&schema); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No schema found with name " + name})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"schema": schema})
}
