package services

import (
	domain "platform-service/internal/core/domain/gas-platform-service"
	"platform-service/internal/core/helper"
	"platform-service/internal/ports"
)

type Service struct {
	platformRepository ports.PlatformRepository
}

func (s *Service) CreateCategory(platform *domain.Category) (*domain.Category, error) {
	if err := helper.Validate(platform); err != nil {
		return nil, err
	}
	return s.platformRepository.CreateCategory(platform)
}

func (s *Service) UpdateCategory(platformReference string, platform *domain.Category) (*domain.Category, error) {
	return s.platformRepository.UpdateCategory(platformReference, platform)
}

func (s *Service) GetCategoryByReference(reference string) ([]*domain.Category, error) {
	return s.platformRepository.GetCategoryByReference(reference)
}

func (s *Service) GetCategoryByName(name string) ([]*domain.Category, error) {
	return s.platformRepository.GetCategoryByName(name)
}

func (s *Service) GetCategoryList(page int64) ([]*domain.Category, error, int64) {
	return s.platformRepository.GetCategoryList(page)
}

func (s *Service) DeleteCategoryByReference(reference string) (interface{}, error) {
	return s.platformRepository.DeleteCategoryByReference(reference)
}

func (s *Service) CreateSubCategory(platform *domain.SubCategory) (*domain.SubCategory, error) {
	if err := helper.Validate(platform); err != nil {
		return nil, err
	}
	return s.platformRepository.CreateSubCategory(platform)
}

func (s *Service) UpdateSubCategory(platformReference string, platform *domain.SubCategory) (*domain.SubCategory, error) {
	return s.platformRepository.UpdateSubCategory(platformReference, platform)
}

func (s *Service) GetSubCategoryByReference(reference string) ([]*domain.SubCategory, error) {
	return s.platformRepository.GetSubCategoryByReference(reference)
}

func (s *Service) GetSubCategoryByName(name string) ([]*domain.SubCategory, error) {
	return s.platformRepository.GetSubCategoryByName(name)
}

func (s *Service) GetSubCategoryList(page int64) ([]*domain.SubCategory, error, int64) {
	return s.platformRepository.GetSubCategoryList(page)
}

func (s *Service) DeleteSubCategoryByReference(reference string) (interface{}, error) {
	return s.platformRepository.DeleteSubCategoryByReference(reference)
}

func (s *Service) CreateState(state *domain.State) (*domain.State, error) {
	if err := helper.Validate(state); err != nil {
		return nil, err
	}
	return s.platformRepository.CreateState(state)
}

func (s *Service) UpdateState(stateReference string, state *domain.State) (*domain.State, error) {
	return s.platformRepository.UpdateState(stateReference, state)
}

func (s *Service) GetStateByReference(reference string) ([]*domain.State, error) {
	return s.platformRepository.GetStateByReference(reference)
}

func (s *Service) GetStateList(page int64) ([]*domain.State, error, int64) {
	return s.platformRepository.GetStateList(page)
}

func (s *Service) DeleteStateByReference(reference string) (interface{}, error) {
	return s.platformRepository.DeleteStateByReference(reference)
}

func New(platformRepository ports.PlatformRepository) *Service {
	return &Service{platformRepository: platformRepository}
}
