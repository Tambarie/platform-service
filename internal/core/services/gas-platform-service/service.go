package services

import (
	domain "platform-service/internal/core/domain/gas-platform-service"
	"platform-service/internal/core/helper"
	"platform-service/internal/ports"
)

type Service struct {
	platformRepository ports.PlatformRepository
}

func (s *Service) GetPlatformPage(page int64) ([]*domain.PlatformService, error, int64) {
	return s.platformRepository.GetPlatformPage(page)
}

func (s *Service) GetCategoryByName(name string) ([]*domain.PlatformService, error) {
	return s.platformRepository.GetCategoryByName(name)
}

func (s *Service) CreatePlatform(platform *domain.PlatformService) (*domain.PlatformService, error) {
	if err := helper.Validate(platform); err != nil {
		return nil, err
	}
	return s.platformRepository.CreatePlatform(platform)
}

func (s *Service) UpdatePlatform(platformReference string, platform *domain.PlatformService) (*domain.PlatformService, error) {
	return s.platformRepository.UpdatePlatform(platformReference, platform)
}

func (s *Service) GetCategoryByReference(reference string) ([]*domain.PlatformService, error) {
	return s.platformRepository.GetCategoryByReference(reference)
}

func (s *Service) DeleteCategoryByReference(reference string) (interface{}, error) {
	return s.platformRepository.DeleteCategoryByReference(reference)
}
func New(platformRepository ports.PlatformRepository) *Service {
	return &Service{platformRepository: platformRepository}
}
