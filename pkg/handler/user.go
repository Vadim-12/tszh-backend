package handler

import (
	"net/http"

	"github.com/Vadim-12/tszh-backend/pkg/handler/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) getMe(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)

	resp, err := h.services.User.GetMe(ctx.Request.Context(), userID)
	if err != nil {
		switch {
		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_SERVER_ERROR", Data: "internal server error"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getAllUsers(ctx *gin.Context) {}

func (h *Handler) getUserById(ctx *gin.Context) {}

func (h *Handler) createUser(ctx *gin.Context) {}

func (h *Handler) updateUser(ctx *gin.Context) {}

func (h *Handler) deleteUser(ctx *gin.Context) {}
