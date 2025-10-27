package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Deps struct {
	Log       *zap.Logger
	Buildings *BuildingHandlers
	Users     *UserHandlers
}

func NewRouter(deps Deps) *gin.Engine {
	rootRouter := gin.New()
	rootRouter.Use(gin.Recovery(), requestLog(deps.Log))

	rootRouter.GET("/healthz", func(c *gin.Context) { c.String(200, "ok") })

	apiRouterGroupV1 := rootRouter.Group("/api/v1")
	deps.Buildings.Register(apiRouterGroupV1)

	return rootRouter
}

func requestLog(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("req",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("dur", time.Since(start)),
		)
	}
}
