// internal/models/store_type.go
package models

import context "context"

type StoreType int

const (
	StoreTypeMongoDB StoreType = iota
	StoreTypePostgres
)

type DataStore interface {
    // get and create schemas
    CreateSchema(ctx context.Context, schema Schema) error
    GetSchema(ctx context.Context, name string) (Schema, error)

    // get and create records
    CreateRecord(ctx context.Context, collection string, record map[string]interface{}) error
    GetSingleRecord(ctx context.Context, collection string, id string) (map[string]interface{}, error)
    GetRecords(ctx context.Context, collection string) ([]map[string]interface{}, error)

    // user handling
    CreateUser(ctx context.Context, user *User) error
    GetUser(ctx context.Context, username string) (User, error)
}
