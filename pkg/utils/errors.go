package utils

import "net/http"


type ErrorCode struct {
    Code    int
    Message string
}

// Определение констант с кодами и сообщениями ошибок
var (
    BadRequest = ErrorCode{
        Code:    http.StatusBadRequest,
        Message: "Bad Request",
    }
    RequestEntityTooLarge = ErrorCode{
        Code:    http.StatusRequestEntityTooLarge,
        Message: "Request body must not exceed 1MB",
    }
    OkRequest = ErrorCode{
        Code:    http.StatusOK,
        Message: "Ok",
    }
	UnAuthorized = ErrorCode{
		Code:	http.StatusUnauthorized,
		Message: "authorized failed",
	} 


)

// Функция для получения указателя на ErrorCode
func GetErrorCode(e ErrorCode) *ErrorCode {
    return &e
}

