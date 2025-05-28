package apiclient

import "fmt"

// APIError 表示API调用中的错误
type APIError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

// NewAPIError 创建一个新的APIError
func NewAPIError(code string, message string, statusCode int) *APIError {
	return &APIError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

// Error 实现error接口
func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("API Error [%s] (HTTP %d): %s", e.Code, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API Error [%s]: %s", e.Code, e.Message)
}
