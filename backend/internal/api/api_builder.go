package api

import (
	"backend/main/internal/config"
	"backend/main/internal/api/handler"
	"backend/main/internal/models"

	"github.com/gin-gonic/gin"
)

type APIBuilder struct {
	Store models.DataStore
	Config *config.Config
}

// Maps routes to handlers
func (api *APIBuilder) SetupRoutes(r *gin.Engine) {
	// setup private API
	private := r.Group("/admin/api")

	private.POST("/schema", handler.CreateSchema(api.Store, api.Config))
	private.GET("/schema/:name", handler.GetSchema(api.Store, api.Config))

	private.POST("/:collection/record", handler.CreateRecord(api.Store, api.Config))
	private.GET("/:collection/record/:id", handler.GetSingleRecord(api.Store, api.Config))
	private.GET("/:collection/records", handler.GetRecords(api.Store, api.Config))

	
	// setup public API
	public := r.Group("/api")

	public.GET("/:collection/record/:id", handler.GetSingleRecord(api.Store, api.Config))
	public.GET("/:collection/records", handler.GetRecords(api.Store, api.Config))
}

func NewAPIBuilder(store models.DataStore, config *config.Config) *APIBuilder {
	return &APIBuilder{Store: store, Config: config}
}
