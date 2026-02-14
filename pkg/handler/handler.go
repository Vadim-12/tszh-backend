package handler

import (
	"time"

	"github.com/Vadim-12/tszh-backend/pkg/handler/middleware"
	"github.com/Vadim-12/tszh-backend/pkg/service"
	"github.com/Vadim-12/tszh-backend/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services  *service.Service
	jwtSigner *utils.JWTSigner
}

func NewHandler(services *service.Service, jwtSigner *utils.JWTSigner) *Handler {
	return &Handler{
		services:  services,
		jwtSigner: jwtSigner,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // если будут cookie/refresh-token в cookie (но пока только мобилка, будем хранить токены SecureStore на клиенте и передавать в Authorization-заголовке)
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.POST("/refresh", h.refresh)
			auth.POST("/logout", h.logout)
		}

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(*h.jwtSigner))
		{
			users := protected.Group("/users")
			{
				users.GET("/me", h.getMe)
				users.GET("", h.getAllUsers)
				users.POST("", h.createUser)
				users.GET("/:id", h.getUserById)
				users.PUT("/:id", h.updateUser)
				users.DELETE("/:id", h.deleteUser)
			}

			buildings := protected.Group("/buildings")
			{
				buildings.GET("", h.getAllBuildings)
				buildings.POST("", h.createBuilding)
				buildings.GET("/:id", h.getBuildingById)
				buildings.PUT("/:id", h.updateBuilding)
				buildings.DELETE("/:id", h.deleteBuilding)
			}

			organizations := protected.Group("/organizations")
			{
				organizations.GET("", h.getAllOrganizations)
				organizations.POST("", h.createOrganization)
				organizations.GET("/:id", h.getOrganizationById)
				organizations.PUT("/:id", h.updateOrganization)
				organizations.DELETE("/:id", h.deleteOrganization)
			}
		}

		api.GET("/health", h.ping)
	}

	return router
}
