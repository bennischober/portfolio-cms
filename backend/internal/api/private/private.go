package private

import (
	"backend/main/internal/config"
	"backend/main/internal/api/private/handlers"
	"backend/main/internal/models"

	"github.com/gin-gonic/gin"
)

type PrivateAPI struct {
    Store models.DataStore
	Config *config.Config
}

// Maps routes to handlers
func (api *PrivateAPI) SetupRoutes(r *gin.Engine) {
	private := r.Group("/admin/api")

	private.POST("/schema", handlers.CreateSchema(api.Store, api.Config))
	private.GET("/schema/:name", handlers.GetSchema(api.Store, api.Config))

	private.POST("/:collection/record", handlers.CreateRecord(api.Store, api.Config))
	private.GET("/:collection/record/:id", handlers.GetSingleRecord(api.Store, api.Config))
	private.GET("/:collection/records", handlers.GetRecords(api.Store, api.Config))
}

func NewPrivateAPI(store models.DataStore, config *config.Config) *PrivateAPI {
	return &PrivateAPI{Store: store, Config: config}
}
