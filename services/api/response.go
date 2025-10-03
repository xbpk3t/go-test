package api

import (
	"time"
)

// APIResponse represents the standard API response format with 6 parameters
type APIResponse struct {
	Code      int         `json:"code"`      // Status code (e.g., 200, 400, 500)
	Data      interface{} `json:"data"`      // Response data
	Msg       string      `json:"msg"`       // Message describing the response
	RequestID string      `json:"requestId"` // Unique identifier for the request
	Success   bool        `json:"success"`   // Indicates if the request was successful
	Timestamp int64       `json:"ts"`        // Unix timestamp of the response
}

// NewSuccessResponse creates a successful API response
func NewSuccessResponse(data interface{}, message string) *APIResponse {
	return &APIResponse{
		Code:      200,
		Data:      data,
		Msg:       message,
		RequestID: generateRequestID(),
		Success:   true,
		Timestamp: time.Now().Unix(),
	}
}

// NewErrorResponse creates an error API response
func NewErrorResponse(code int, message string) *APIResponse {
	return &APIResponse{
		Code:      code,
		Data:      nil,
		Msg:       message,
		RequestID: generateRequestID(),
		Success:   false,
		Timestamp: time.Now().Unix(),
	}
}

// generateRequestID generates a unique request ID (simplified implementation)
func generateRequestID() string {
	return time.Now().Format("20060102150405")
}