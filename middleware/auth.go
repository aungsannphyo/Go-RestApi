package middleware

import (
	"net/http"
	"strings"

	"github.com/aungsannphyo/go-restapi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	header := context.Request.Header.Get("Authorization")

	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": " Authorization header missing or malformed"})
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
