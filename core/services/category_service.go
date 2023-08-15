package services

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"
)

type categoryService struct {
	catRepo ports.CategoryRepository
}

func NewCategoryService(catRepo ports.CategoryRepository) ports.CategoryService {
	return categoryService{catRepo}
}

func (s categoryService) Create(ctg *model_io.Category) error {
	category := model_gorm.Category{
		Name: ctg.Name,
	}

	err := s.catRepo.Create(&category)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s categoryService) Update(ctg *model_io.Category, cId *int) error {
	category := model_gorm.Category{
		Name: ctg.Name,
	}

	err := s.catRepo.Update(&category, cId)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s categoryService) Delete(cId *int) error {
	err := s.catRepo.Delete(cId)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s categoryService) GetCategories() (*[]model_io.Category, error) {
	categories, err := s.catRepo.GetCategories()
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return categories, nil
}

func (s categoryService) GetCategoryById(cId *int) (*model_io.Category, error) {
	category, err := s.catRepo.GetCategoryById(cId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return category, nil
}
