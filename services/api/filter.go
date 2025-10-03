package api

import (
	"fmt"
	"net/http"
	"strings"
)

// FilterParams holds the parsed filter parameters
type FilterParams struct {
	Filters map[string]string
	Fields  []string
	Sort    []string
}

// ParseFilterParams parses filter parameters from the request
func ParseFilterParams(r *http.Request) *FilterParams {
	params := &FilterParams{
		Filters: make(map[string]string),
		Fields:  []string{},
		Sort:    []string{},
	}

	// Parse filter parameters
	for key, values := range r.URL.Query() {
		if key == "filter" {
			for _, value := range values {
				// Parse filter format: field:operator:value
				parts := strings.Split(value, ":")
				if len(parts) == 3 {
					filterKey := fmt.Sprintf("%s:%s", parts[0], parts[1])
					params.Filters[filterKey] = parts[2]
				}
			}
		} else if key == "fields" {
			for _, value := range values {
				params.Fields = strings.Split(value, ",")
			}
		} else if key == "sort" {
			for _, value := range values {
				params.Sort = strings.Split(value, ",")
			}
		}
	}

	return params
}

// ApplyFilters applies filters to a GORM query
func ApplyFilters(query interface{}, filterParams *FilterParams) interface{} {
	// This is a simplified implementation
	// In a real application, you would apply the filters to a GORM query

	// Example of how filtering might work:
	// if value, exists := filterParams.Filters["name:eq"]; exists {
	//     query = query.Where("name = ?", value)
	// }

	// For sorting:
	// for _, sortField := range filterParams.Sort {
	//     if strings.HasPrefix(sortField, "-") {
	//         query = query.Order(fmt.Sprintf("%s DESC", strings.TrimPrefix(sortField, "-")))
	//     } else {
	//         query = query.Order(fmt.Sprintf("%s ASC", sortField))
	//     }
	// }

	return query
}

// SelectFields selects only specified fields from the data
func SelectFields(data interface{}, fields []string) interface{} {
	if len(fields) == 0 {
		return data
	}

	// This is a simplified implementation
	// In a real application, you would select only the specified fields
	// This could be done at the database level for efficiency

	// For now, we'll just return the data as-is
	// A full implementation would use reflection to create a new struct
	// with only the specified fields

	return data
}