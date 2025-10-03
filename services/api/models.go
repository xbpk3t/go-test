package api

import (
	"gorm.io/gorm"
)

// SampleModel represents a sample data model for testing
type SampleModel struct {
	gorm.Model
	Name  string `json:"name"`
	Value string `json:"value"`
}