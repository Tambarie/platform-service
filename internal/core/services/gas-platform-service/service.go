package services

import "platform-service/internal/ports"

type Service struct {
	platformRepository ports.PlatformRepository
}

func New(platformRepository ports.PlatformRepository) *Service {
	return &Service{platformRepository: platformRepository}
}
