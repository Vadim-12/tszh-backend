package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomError struct {
	Error CustomErrorBody `json:"error"`
}

type CustomErrorBody struct {
	Code string `json:"code"`
	Data any    `json:"data,omitempty"`
}

func NewResponseError(ctx *gin.Context, statusCode int, err CustomErrorBody) {
	logrus.Error(err.Data)
	ctx.AbortWithStatusJSON(statusCode, CustomErrorBody{Code: err.Code, Data: err.Data})
}
