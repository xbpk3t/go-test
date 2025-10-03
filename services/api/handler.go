package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// HandlerFunc defines the API handler function signature
type HandlerFunc func(http.ResponseWriter, *http.Request) *APIResponse

// Handle wraps a handler function with common functionality
func Handle(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set common headers
		w.Header().Set("Content-Type", "application/json")

		// Process the request
		response := handler(w, r)

		// Send the response
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

// ExampleHandler demonstrates how to use the API components
func ExampleHandler(w http.ResponseWriter, r *http.Request) *APIResponse {
	// Parse filter parameters
	filterParams := ParseFilterParams(r)

	// For demonstration, let's create a sample record if none exists
	var count int64
	DB.Model(&SampleModel{}).Count(&count)
	if count == 0 {
		// Create a sample record
		sample := SampleModel{
			Name:  "Sample Item",
			Value: "Sample Value",
		}
		DB.Create(&sample)
	}

	// Query the database
	var samples []SampleModel
	query := DB.Model(&SampleModel{})

	// Apply filters
	for filterKey, filterValue := range filterParams.Filters {
		parts := strings.Split(filterKey, ":")
		if len(parts) == 2 {
			field := parts[0]
			operator := parts[1]

			switch operator {
			case "eq":
				query = query.Where(field+" = ?", filterValue)
			case "ne":
				query = query.Where(field+" != ?", filterValue)
			case "gt":
				query = query.Where(field+" > ?", filterValue)
			case "lt":
				query = query.Where(field+" < ?", filterValue)
			}
		}
	}

	// Apply sorting
	for _, sortField := range filterParams.Sort {
		if strings.HasPrefix(sortField, "-") {
			query = query.Order(strings.TrimPrefix(sortField, "-") + " DESC")
		} else {
			query = query.Order(sortField + " ASC")
		}
	}

	// Execute query
	if err := query.Find(&samples).Error; err != nil {
		return NewErrorResponse(500, "Failed to query database: "+err.Error())
	}

	// For field selection, we'll need to implement a more complex solution
	// For now, we'll return the full data
	result := samples

	return NewSuccessResponse(result, "Successfully retrieved data")
}

// CreateSampleHandler creates a new sample record
func CreateSampleHandler(w http.ResponseWriter, r *http.Request) *APIResponse {
	if r.Method != http.MethodPost {
		return NewErrorResponse(405, "Method not allowed")
	}

	// Parse form data
	name := r.FormValue("name")
	value := r.FormValue("value")

	if name == "" || value == "" {
		return NewErrorResponse(400, "Name and value are required")
	}

	// Create new sample record
	sample := SampleModel{
		Name:  name,
		Value: value,
	}

	if err := DB.Create(&sample).Error; err != nil {
		return NewErrorResponse(500, "Failed to create record: "+err.Error())
	}

	return NewSuccessResponse(sample, "Successfully created record")
}

// GetSampleHandler retrieves a specific sample record by ID
func GetSampleHandler(w http.ResponseWriter, r *http.Request) *APIResponse {
	// Extract ID from URL path
	idStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return NewErrorResponse(400, "Invalid ID format")
	}

	var sample SampleModel
	if err := DB.First(&sample, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewErrorResponse(404, "Record not found")
		}
		return NewErrorResponse(500, "Database error: "+err.Error())
	}

	return NewSuccessResponse(sample, "Successfully retrieved record")
}