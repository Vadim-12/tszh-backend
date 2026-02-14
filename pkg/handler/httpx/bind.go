package httpx

import (
	"log"
	"net/http"

	"github.com/Vadim-12/tszh-backend/pkg/handler/validation"
	"github.com/gin-gonic/gin"
)

func BindAndValidate[T any](ctx *gin.Context, payload *T) bool {
	if err := ctx.BindJSON(payload); err != nil {
		NewResponseError(ctx, http.StatusBadRequest, CustomErrorBody{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
		return false
	}

	if err := validation.Validate(payload); err != nil {
		log.Println("VALIDATION ERROR", err)
		NewResponseError(ctx, http.StatusBadRequest, CustomErrorBody{
			Code: "BAD_REQUEST",
			Data: validation.ToValidationErrors(err),
		})
		return false
	}

	return true
}
