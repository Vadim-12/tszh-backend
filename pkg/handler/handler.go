package handler

import (
	"github.com/Vadim-12/tszh-backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.POST("/refresh", h.refresh)
			auth.POST("/logout", h.logout)
		}

		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.POST("/", h.createUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}

		buildings := api.Group("/buildings")
		{
			buildings.GET("/", h.getAllBuildings)
			buildings.GET("/:id", h.getBuildingById)
			buildings.POST("/", h.createBuilding)
			buildings.PUT("/:id", h.updateBuilding)
			buildings.DELETE("/:id", h.deleteBuilding)
		}

		organizations := api.Group("/organizations")
		{
			organizations.GET("/", h.getAllOrganizations)
			organizations.GET("/:id", h.getOrganizationById)
			organizations.POST("/", h.createOrganization)
			organizations.PUT("/:id", h.updateOrganization)
			organizations.DELETE("/:id", h.deleteOrganization)
		}

		api.GET("/health", h.ping)
	}

	return router
}
