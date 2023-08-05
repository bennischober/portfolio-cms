// internal/data/provider/mongodb/store.go
package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cms/main/internal/models"
)

type MongoDBStore struct {
	client *mongo.Client
	dbName string
}

// hardcoded schema collection name
const (
	SCHEMA_COLLECTION = "schema"
)

func (s *MongoDBStore) EnsureIndexes(ctx context.Context) error {
	// create unique index on name field
	collection := s.client.Database(s.dbName).Collection(SCHEMA_COLLECTION)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1, // index in ascending order
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	return nil
}

// CreateSchema creates a new schema in the database
func (s *MongoDBStore) CreateSchema(ctx context.Context, schema models.Schema) error {
	collection := s.client.Database(s.dbName).Collection(SCHEMA_COLLECTION)

	if _, err := collection.InsertOne(ctx, schema); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	return nil
}

func (s *MongoDBStore) GetSchema(ctx context.Context, name string) (models.Schema, error) {
	collection := s.client.Database(s.dbName).Collection(SCHEMA_COLLECTION)

	var schema models.Schema
	if err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&schema); err != nil {
		if err == mongo.ErrNoDocuments {
			return schema, fmt.Errorf("schema not found: %w", err)
		} else {
			return schema, fmt.Errorf("failed to get schema: %w", err)
		}
	}

	return schema, nil
}

func (s *MongoDBStore) CreateRecord(ctx context.Context, schemaName string, data map[string]interface{}) error {
	collection := s.client.Database(s.dbName).Collection(schemaName)

    // check, if schema exists
    schema, err := s.GetSchema(ctx, schemaName)
    if err != nil {
        return fmt.Errorf("failed to get schema: %w", err)
    }

    // validate data
    if err := validateData(schema, data); err != nil {
        return fmt.Errorf("failed to validate data: %w", err)
    }

    // schema exists and data is valid, create record
	if _, err := collection.InsertOne(ctx, data); err != nil {
		return fmt.Errorf("failed to create record: %w", err)
	}

	return nil
}

func (s *MongoDBStore) GetSingleRecord(ctx context.Context, schemaName string, id string) (map[string]interface{}, error) {
	collection := s.client.Database(s.dbName).Collection(schemaName)

    // create object id from string
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("failed to create object id: %w", err)
    }

	var record map[string]interface{}
	if err := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&record); err != nil {
		if err == mongo.ErrNoDocuments {
			return record, fmt.Errorf("record not found: %w", err)
		} else {
			return record, fmt.Errorf("failed to get record: %w", err)
		}
	}

	return record, nil
}

func (s *MongoDBStore) GetRecords(ctx context.Context, schemaName string) ([]map[string]interface{}, error) {
	collection := s.client.Database(s.dbName).Collection(schemaName)

	var records []map[string]interface{}
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return records, fmt.Errorf("failed to get records: %w", err)
	}
	if err := cursor.All(context.Background(), &records); err != nil {
		return records, fmt.Errorf("failed to get records: %w", err)
	}

	return records, nil
}

func validateData(schema models.Schema, data map[string]interface{}) error {
	for _, field := range schema.Fields {
		// Check if field exists in data
		if _, ok := data[field.Name]; !ok {
			return fmt.Errorf("missing field %s", field.Name)
		}
	}
	return nil
}

func NewMongoDBStore(ctx context.Context, uri string, dbName string) (*MongoDBStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Create new store
	store := &MongoDBStore{client: client, dbName: dbName}

	// Ensure, that all indexes are created
	if err := store.EnsureIndexes(ctx); err != nil {
		return nil, err
	}

	return store, nil
}
