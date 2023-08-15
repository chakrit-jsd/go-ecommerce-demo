package repositories

import (
	"errors"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepositoryDB(db *gorm.DB) ports.CategoryRepository {
	return categoryRepository{db}
}

func (r categoryRepository) Create(ctg *model_gorm.Category) error {
	err := r.db.Model(&model_gorm.Category{}).Create(&ctg).Error
	if err != nil {
		return err
	}

	return nil
}

func (r categoryRepository) Update(ctg *model_gorm.Category, cId *int) error {
	res := r.db.Model(&model_gorm.Category{}).Updates(&ctg)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("update fail id not found")
	}

	return nil
}

func (r categoryRepository) Delete(cId *int) error {
	res := r.db.Model(&model_gorm.Category{}).Delete(cId)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("delete fail id not found")
	}

	return nil
}

func (r categoryRepository) GetCategories() (*[]model_io.Category, error) {
	categories := []model_io.Category{}
	err := r.db.Model(&model_gorm.Category{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (r categoryRepository) GetCategoryById(cId *int) (*model_io.Category, error) {
	category := model_io.Category{}
	err := r.db.Model(&model_gorm.Category{}).Where("id = ?", cId).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// func (r categoryRepository) GetCategoryByName(cName *string) (*model_io.Category, error) {
// 	return nil, nil
// }
