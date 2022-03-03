package api

import (
	"github.com/gin-gonic/gin"
	"platform-service/internal/ports"
)

type HTTPHandler struct {
	platformService ports.PlatformService
}

func NewHTTPHandler(platformService ports.PlatformService) *HTTPHandler {
	return &HTTPHandler{
		platformService: platformService,
	}
}

func (h *HTTPHandler) CreatePlatform() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
