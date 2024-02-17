package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var (
	BAD_REQUEST_STATUS_CODE = http.StatusBadRequest
	CREATED_REQUEST_STATUS_CODE = http.StatusCreated
	SUCCESS_REQUEST_STATUS_CODE = http.StatusOK
	ERROR_REQUEST_STATUS_CODE = http.StatusInternalServerError
	NOT_FOUND_REQUEST_STATUS_CODE = http.StatusNotFound
)

type ResponseCommonParameters struct {
	Status_Code int `json:"status_code"`
	Message string `json:"message"`
}

func ValidationErrorResponseH(err error) gin.H {
	return gin.H{
		"status_code": BAD_REQUEST_STATUS_CODE,
		"message": "Validation failed",
		"error": err.Error(),
	}
}

func ErrorResponseH(err error) gin.H {
	return gin.H{
		"status_code": ERROR_REQUEST_STATUS_CODE,
		"message": "Something went wrong, please try again",
		"error": err.Error(),
	}
}

func SuccessResponse(data any) gin.H {
	return gin.H{
		"status_code": SUCCESS_REQUEST_STATUS_CODE,
		"message": "Success",
		"data": data,
	}
}

func NotFoundErrorResponseH(err error) gin.H {
	return gin.H{
		"status_code": NOT_FOUND_REQUEST_STATUS_CODE,
		"message": "Record not found",
		"error": err.Error(),
	}
}