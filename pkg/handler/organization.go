package handler

import (
	"errors"
	"net/http"

	appErrors "github.com/Vadim-12/tszh-backend/pkg/errors"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/handler/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) getAllOrganizations(ctx *gin.Context) {
	resp, err := h.services.Organization.GetAll(ctx.Request.Context())
	if err != nil {
		switch {
		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_SERVER_ERROR"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getOrganizationById(ctx *gin.Context) {
	orgId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httpx.NewResponseError(ctx, http.StatusBadRequest, httpx.CustomErrorBody{Code: "INVALID_ORGANIZATION_ID"})
		return
	}

	resp, err := h.services.Organization.GetByID(ctx.Request.Context(), orgId)
	if err != nil {
		switch err {
		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_SERVER_ERROR"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) createOrganization(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)

	var input dto.CreateOrganizationPayload
	if good := httpx.BindAndValidate(ctx, &input); !good {
		return
	}

	resp, err := h.services.Organization.Create(ctx.Request.Context(), &input, userID)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrINNAlreadyExists):
			httpx.NewResponseError(ctx, http.StatusConflict, httpx.CustomErrorBody{Code: "INN_ALREADY_EXISTS"})
			return

		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_SERVER_ERROR"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateOrganization(ctx *gin.Context) {}

func (h *Handler) deleteOrganization(ctx *gin.Context) {}
