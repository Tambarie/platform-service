package api

import (
	"gas-platform-service/internal/ports"
)

type HTTPHandler struct {
	platformService ports.PlatformService
}

func NewHTTPHandler(platformService ports.PlatformService) *HTTPHandler {
	return &HTTPHandler{
		platformService: platformService,
	}
}
