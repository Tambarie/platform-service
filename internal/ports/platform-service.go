package ports

import domain "platform-service/internal/core/domain/gas-platform-service"

type PlatformService interface {
	CreateCategory(platform *domain.Category) (*domain.Category, error)
	UpdateCategory(platformReference string, platform *domain.Category) (*domain.Category, error)
	GetCategoryByReference(reference string) ([]*domain.Category, error)
	GetCategoryByName(reference string) ([]*domain.Category, error)
	GetCategoryList(page int64) ([]*domain.Category, error, int64)
	DeleteCategoryByReference(reference string) (interface{}, error)

	CreateSubCategory(platform *domain.SubCategory) (*domain.SubCategory, error)
	UpdateSubCategory(platformReference string, platform *domain.SubCategory) (*domain.SubCategory, error)
	GetSubCategoryByReference(reference string) ([]*domain.SubCategory, error)
	GetSubCategoryByName(reference string) ([]*domain.SubCategory, error)
	GetSubCategoryList(page int64) ([]*domain.SubCategory, error, int64)
	DeleteSubCategoryByReference(reference string) (interface{}, error)
}
