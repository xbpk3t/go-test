package api

// ErrorCode represents a standardized error code
type ErrorCode int

// Standard error codes
const (
	// 2xx Success
	Success ErrorCode = 200

	// 4xx Client Errors
	BadRequest          ErrorCode = 400
	Unauthorized        ErrorCode = 401
	Forbidden           ErrorCode = 403
	NotFound            ErrorCode = 404
	MethodNotAllowed    ErrorCode = 405
	Conflict            ErrorCode = 409
	UnprocessableEntity ErrorCode = 422

	// 5xx Server Errors
	InternalServerError ErrorCode = 500
	NotImplemented      ErrorCode = 501
	BadGateway          ErrorCode = 502
	ServiceUnavailable  ErrorCode = 503
	GatewayTimeout      ErrorCode = 504
)

// ErrorInfo holds detailed information about an error
type ErrorInfo struct {
	Code    ErrorCode
	Message string
	Details map[string]interface{}
}

// ErrorMessages maps error codes to default messages
var ErrorMessages = map[ErrorCode]string{
	Success:             "Success",
	BadRequest:          "Bad Request",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not Found",
	MethodNotAllowed:    "Method Not Allowed",
	Conflict:            "Conflict",
	UnprocessableEntity: "Unprocessable Entity",
	InternalServerError: "Internal Server Error",
	NotImplemented:      "Not Implemented",
	BadGateway:          "Bad Gateway",
	ServiceUnavailable:  "Service Unavailable",
	GatewayTimeout:      "Gateway Timeout",
}

// NewErrorInfo creates a new ErrorInfo with default message
func NewErrorInfo(code ErrorCode, details map[string]interface{}) *ErrorInfo {
	message, exists := ErrorMessages[code]
	if !exists {
		message = "Unknown Error"
	}

	return &ErrorInfo{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// NewCustomErrorInfo creates a new ErrorInfo with custom message
func NewCustomErrorInfo(code ErrorCode, message string, details map[string]interface{}) *ErrorInfo {
	return &ErrorInfo{
		Code:    code,
		Message: message,
		Details: details,
	}
}