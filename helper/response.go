package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type apiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func APIResponse(code int, message string, status string, data interface{}) apiResponse {
	jsonResponse := apiResponse{
		Code:    code,
		Message: message,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors

}

func ErrorValidation(c *gin.Context, err error, message string) {
	errors := FormatValidationError(err)
	errorMessage := gin.H{"errors": errors}

	response := APIResponse(http.StatusBadRequest, message, "error", errorMessage)
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func ErrorHandling(c *gin.Context, err error, message string) {
	var errors []string
	errors = append(errors, err.Error())
	errorMessage := gin.H{"errors": errors}
	response := APIResponse(http.StatusUnprocessableEntity, message, "error", errorMessage)
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
}

func SuccessHandling(c *gin.Context, formatter interface{}, message string) {
	response := APIResponse(http.StatusOK, message, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func AuthorizationHandling(c *gin.Context) {
	var errors []string
	errors = append(errors, "You're Not Authorized to do this")
	errorMessage := gin.H{"errors": errors}

	response := APIResponse(http.StatusUnauthorized, "Unauthorized", "error", errorMessage)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)

}
