package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Vadim-12/tszh-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSigner utils.JWTSigner) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		log.Println("authHeader", authHeader)
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing auth header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}

		userID, err := jwtSigner.ParseAccess(parts[1])
		log.Println("userID from access token:", userID)
		if err != nil {
			log.Println("ParseAccess error:", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		log.Println("FROM TOKEN: user_id =", userID)
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
