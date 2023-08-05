// internal/models/store_type.go
package models

type Schema struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name     string      `json:"name"`
	DataType string      `json:"data_type"`
}