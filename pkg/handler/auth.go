package handler

import (
	"context"
	"errors"
	"net/http"

	appErrors "github.com/Vadim-12/tszh-backend/pkg/errors"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input dto.SignUpRequestDto
	if err := ctx.BindJSON(&input); err != nil {
		newResponseError(ctx, http.StatusBadRequest, CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.SignUp(ctx.Request.Context(), &input)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrUserWithPhoneNumberAlreadyExists):
			newResponseError(ctx, http.StatusConflict, CustomErrorBody{Code: "PHONE_NUMBER_ALREADY_EXISTS", Data: appErrors.ErrUserWithPhoneNumberAlreadyExists.Error()})
			return

		case errors.Is(err, appErrors.ErrUserWithEmailAlreadyExists):
			newResponseError(ctx, http.StatusConflict, CustomErrorBody{Code: "EMAIL_ALREADY_EXISTS", Data: appErrors.ErrUserWithEmailAlreadyExists.Error()})
			return

		case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
			newResponseError(ctx, http.StatusGatewayTimeout, CustomErrorBody{Code: "REQUEST_TIMEOUT", Data: "request timeout"})
			return

		default:
			newResponseError(ctx, http.StatusInternalServerError, CustomErrorBody{Code: "INTERNAL_SERVER_ERROR", Data: "internal server error"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input dto.SignInRequestDto
	if err := ctx.BindJSON(&input); err != nil {
		newResponseError(ctx, http.StatusBadRequest, CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.SignIn(ctx.Request.Context(), &input)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidCredentials):
			newResponseError(ctx, http.StatusUnauthorized, CustomErrorBody{Code: "WRONG_CREDS", Data: "wrong phone number of password"})
			return

		case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
			newResponseError(ctx, http.StatusGatewayTimeout, CustomErrorBody{Code: "REQUEST_TIMEOUT", Data: "request timeout"})
			return

		default:
			newResponseError(ctx, http.StatusInternalServerError, CustomErrorBody{Code: "INTERNAL_ERROR", Data: "internal server error"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) refresh(ctx *gin.Context) {
	var input dto.RefreshRequestDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newResponseError(ctx, http.StatusBadRequest, CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.Refresh(ctx.Request.Context(), &input)
	if err != nil {
		if errors.Is(err, appErrors.ErrRefreshInvalid) {
			newResponseError(ctx, http.StatusBadRequest, CustomErrorBody{Code: "REFRESH_INVALID", Data: err.Error()})
			return
		}
		newResponseError(ctx, http.StatusInternalServerError, CustomErrorBody{Code: "INTERNAL_ERROR", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) logout(ctx *gin.Context) {
	var input dto.LogoutRequestDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newResponseError(ctx, http.StatusBadRequest, CustomErrorBody{Code: "BAD_REQUEST", Data: err.Error()})
		return
	}

	resp, err := h.services.Authorization.Logout(ctx.Request.Context(), &input)
	if err != nil {
		newResponseError(ctx, http.StatusInternalServerError, CustomErrorBody{Code: "INTERNAL_ERROR", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
