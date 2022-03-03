package ports

import domain "platform-service/internal/core/domain/gas-platform-service"

type PlatformRepository interface {
	CreatePlatform(platform *domain.PlatformService) (*domain.PlatformService, error)
	UpdatePlatform(platformReference string, platform *domain.PlatformService) (*domain.PlatformService, error)
	GetCategoryByReference(reference string) ([]*domain.PlatformService, error)
	GetCategoryByName(reference string) ([]*domain.PlatformService, error)
	GetPlatformPage(page int64) ([]*domain.PlatformService, error, int64)
	DeleteCategoryByReference(reference string) (interface{}, error)
}
