package handler

import (
	"context"
	"errors"
	"net/http"

	appErrors "github.com/Vadim-12/tszh-backend/pkg/errors"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/handler/httpx"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input dto.SignUpRequestDto
	if good := httpx.BindAndValidate(ctx, &input); !good {
		return
	}

	resp, err := h.services.Authorization.SignUp(ctx.Request.Context(), &input)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrUserWithPhoneNumberAlreadyExists):
			httpx.NewResponseError(ctx, http.StatusConflict, httpx.CustomErrorBody{Code: "PHONE_NUMBER_ALREADY_EXISTS", Data: appErrors.ErrUserWithPhoneNumberAlreadyExists.Error()})
			return

		case errors.Is(err, appErrors.ErrUserWithEmailAlreadyExists):
			httpx.NewResponseError(ctx, http.StatusConflict, httpx.CustomErrorBody{Code: "EMAIL_ALREADY_EXISTS", Data: appErrors.ErrUserWithEmailAlreadyExists.Error()})
			return

		case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
			httpx.NewResponseError(ctx, http.StatusGatewayTimeout, httpx.CustomErrorBody{Code: "REQUEST_TIMEOUT", Data: err.Error()})
			return

		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_SERVER_ERROR", Data: err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input dto.SignInRequestDto
	if err := ctx.BindJSON(&input); err != nil {
		httpx.NewResponseError(ctx, http.StatusBadRequest, httpx.CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.SignIn(ctx.Request.Context(), &input)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidCredentials):
			httpx.NewResponseError(ctx, http.StatusUnauthorized, httpx.CustomErrorBody{Code: "WRONG_CREDS", Data: "wrong phone number of password"})
			return

		case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
			httpx.NewResponseError(ctx, http.StatusGatewayTimeout, httpx.CustomErrorBody{Code: "REQUEST_TIMEOUT", Data: err.Error()})
			return

		default:
			httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_ERROR", Data: err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) refresh(ctx *gin.Context) {
	var input dto.RefreshRequestDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httpx.NewResponseError(ctx, http.StatusBadRequest, httpx.CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.Refresh(ctx.Request.Context(), &input)
	if err != nil {
		if errors.Is(err, appErrors.ErrRefreshInvalid) {
			httpx.NewResponseError(ctx, http.StatusBadRequest, httpx.CustomErrorBody{Code: "REFRESH_INVALID", Data: err.Error()})
			return
		}
		httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_ERROR", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) logout(ctx *gin.Context) {
	var input dto.LogoutRequestDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httpx.NewResponseError(ctx, http.StatusBadRequest, httpx.CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.Logout(ctx.Request.Context(), &input)
	if err != nil {
		httpx.NewResponseError(ctx, http.StatusInternalServerError, httpx.CustomErrorBody{Code: "INTERNAL_ERROR", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
