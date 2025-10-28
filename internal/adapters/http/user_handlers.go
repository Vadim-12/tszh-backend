package http

import (
	"strconv"

	"github.com/Vadim-12/tszh-backend/internal/domains/users"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct{ Service *users.Service }

func (h *UserHandlers) Register(g *gin.RouterGroup) {
	g.GET("/users/:id", h.getById)
	g.GET("/users", h.getAll)
	g.POST("/users", h.create)
	g.PUT("/users/:id", h.updateOne)
	g.DELETE("/users/:id", h.deleteOne)
}

func (h *UserHandlers) getById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := h.Service.Get(ctx, userId)

	if err != nil {
		if err == users.ErrNotFound {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}

func (h *UserHandlers) getAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	users, err := h.Service.List(ctx, limit, offset)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (h *UserHandlers) create(ctx *gin.Context) {
	var creationDto struct {
		FullName    string `json:"full_name"    binding:"required"`
		Email       string `json:"email"        binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&creationDto); err != nil {
		ctx.JSON(400, gin.H{"error": "Not valid data format for user creation"})
	}

	in := users.User{
		FullName:    creationDto.FullName,
		Email:       creationDto.Email,
		PhoneNumber: creationDto.PhoneNumber,
	}

	createdUser, err := h.Service.Create(ctx, in)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, createdUser)
}

func (h *UserHandlers) updateOne(ctx *gin.Context) {
	userId := ctx.Param("id")
	var updateDto struct {
		FullName    string `json:"full_name"    binding:"required"`
		Email       string `json:"email"        binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&updateDto); err != nil {
		ctx.JSON(400, gin.H{"error": "Not valid data format for update user"})
		return
	}

	in := users.User{
		FullName:    updateDto.FullName,
		Email:       updateDto.Email,
		PhoneNumber: updateDto.PhoneNumber,
	}

	out, err := h.Service.Update(ctx, userId, in)
	if err != nil {
		if err == users.ErrNotFound {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, out)
}

func (h *UserHandlers) deleteOne(ctx *gin.Context) {
	userId := ctx.Param("id")
	if err := h.Service.Delete(ctx, userId); err != nil {
		if err == users.ErrNotFound {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(204)
}
