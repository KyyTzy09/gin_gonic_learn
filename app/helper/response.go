package helper

import (
	"gin-01/app/dto/response"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, response.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, response.BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func ValidationErrorResponse(c *gin.Context, err interface{}) {
	c.JSON(http.StatusBadRequest, response.BaseResponse{
		Success: false,
		Message: "Validation error",
		Error:   err,
	})
}

func PaginationResponse(c *gin.Context, message string, data interface{}, page, pageSize int, totalRows int64) {
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))

	c.JSON(http.StatusOK, response.PaginationResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: response.Pagination{
			Page:       page,
			PageSize:   pageSize,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
	})
}