// internal/data/datastore_factory.go
package data

import (
	"context"
	"fmt"
	"strings"

	"backend/main/internal/models"
	"backend/main/internal/data/provider/mongodb"
	//"backend/main/internal/data/provider/postgres"

	"backend/main/internal/config"
)

func NewDataStore(ctx context.Context, config *config.Config) (models.DataStore, error) {
	storeType := storeTypeFromString(config.Database.Type)
	connString := config.Database.Connect

	switch storeType {
	case models.StoreTypeMongoDB:
		return mongodb.NewMongoDBStore(ctx, connString, config.Database.Name)
	case models.StoreTypePostgres:
		// throw not implemented error
		return nil, fmt.Errorf("provider error: Postgres store not implemented yet")
		//return postgres.NewPostgresStore(ctx, connString)
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Invalid store type: %v", storeType))
	}
}

func storeTypeFromString(storeTypeStr string) (models.StoreType) {
	switch strings.ToLower(storeTypeStr) {
	case "mongodb":
		return models.StoreTypeMongoDB
	case "postgres":
		return models.StoreTypePostgres
	default:
		return models.StoreTypeMongoDB
	}
}
