package http

import (
	"strconv"

	"github.com/Vadim-12/tszh-backend/internal/domains/buildings"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type BuildingHandlers struct{ Service *buildings.Service }

func (h *BuildingHandlers) Register(g *gin.RouterGroup) {
	g.GET("/buildings/:id", h.get)
	g.GET("/buildings", h.getAll)
	g.POST("/buildings", h.create)
	g.PUT("/buildings/:id", h.update)
	g.DELETE("/buildings/:id", h.deleteOne)
}

func (h *BuildingHandlers) get(c *gin.Context) {
	buildingId := c.Param("id")
	out, err := h.Service.Get(c, buildingId)

	if err != nil {
		if err == buildings.ErrNotFound {
			c.JSON(404, gin.H{"error": "Building not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, out)
}

func (h *BuildingHandlers) getAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	buildings, err := h.Service.List(ctx, limit, offset)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, buildings)
}

func (h *BuildingHandlers) create(ctx *gin.Context) {
	var creationDto struct {
		Number    string `json:"number"     binding:"required"`
		Floor     int    `json:"floor"      binding:"required"`
		OwnerName string `json:"owner_name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&creationDto); err != nil {
		ctx.JSON(400, gin.H{"error": "Not valid data format for building creation"})
		return
	}

	in := buildings.Building{
		ID:        uuid.New(),
		Number:    creationDto.Number,
		Floor:     creationDto.Floor,
		OwnerName: creationDto.OwnerName,
	}

	createdBuilding, err := h.Service.Create(ctx, in)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, createdBuilding)
}

func (h *BuildingHandlers) update(ctx *gin.Context) {
	buildingId := ctx.Param("id")
	var updateDto struct {
		Number    string `json:"number"     binding:"required"`
		Floor     int    `json:"floor"      binding:"required"`
		OwnerName string `json:"owner_name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&updateDto); err != nil {
		ctx.JSON(400, gin.H{"error": "Not valid data format for update building"})
		return
	}

	in := buildings.Building{
		Number:    updateDto.Number,
		Floor:     updateDto.Floor,
		OwnerName: updateDto.OwnerName,
	}

	out, err := h.Service.Update(ctx, buildingId, in)
	if err != nil {
		if err == buildings.ErrNotFound {
			ctx.JSON(404, gin.H{"error": "Building not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, out)
}

func (h *BuildingHandlers) deleteOne(ctx *gin.Context) {
	buildingId := ctx.Param("id")
	if err := h.Service.Delete(ctx, buildingId); err != nil {
		if err == buildings.ErrNotFound {
			ctx.JSON(404, gin.H{"error": "Building not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(204)
}
