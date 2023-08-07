package api

import (
	"backend/main/internal/api/handler"
	"backend/main/internal/auth"
	"backend/main/internal/config"
	"backend/main/internal/models"

	"github.com/gin-gonic/gin"
)

type APIBuilder struct {
	Store      models.DataStore
	Config     *config.Config
	Auth       auth.IAuthService
	Middleware auth.IAuthMiddleware
}

// Maps routes to handlers
func (api *APIBuilder) SetupRoutes(r *gin.Engine) {
	// setup private API
	private := r.Group("/admin/api")
	private.Use(api.Middleware.Authenticate())

	private.POST("/schema", handler.CreateSchema(api.Store, api.Config))
	private.GET("/schema/:name", handler.GetSchema(api.Store, api.Config))

	private.POST("/:collection/record", handler.CreateRecord(api.Store, api.Config))

	// setup public/shared API
	public := r.Group("/api")

	public.GET("/:collection/record/:id", handler.GetSingleRecord(api.Store, api.Config))
	public.GET("/:collection/records", handler.GetRecords(api.Store, api.Config))

	// authentication
	authentication := public.Group("/auth")
	authentication.POST("/register", handler.RegisterUser(api.Auth, api.Config))
	authentication.POST("/login", handler.LoginUser(api.Auth, api.Config))
}

func NewAPIBuilder(store models.DataStore, config *config.Config, auth auth.IAuthService, middleware auth.IAuthMiddleware) *APIBuilder {
	return &APIBuilder{Store: store, Config: config, Auth: auth, Middleware: middleware}
}
