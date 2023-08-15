package repositories

import (
	"errors"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	return productRepository{db}
}

func (r productRepository) Create(pd *model_gorm.Product) error {
	err := r.db.Model(&model_gorm.Product{}).Create(&pd).Error
	if err != nil {
		return err
	}
	return nil
}

func (r productRepository) Update(pd *model_gorm.Product, pId *int) error {
	res := r.db.Model(&model_gorm.Product{}).Where("id = ?", pId).Updates(&pd)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("update fail id not found")
	}

	return nil
}

func (r productRepository) UpdateStock(qty, pId *int) error {
	res := r.db.Model(&model_gorm.Product{}).Where("id = ?", pId).Update("stock", gorm.Expr("stock + ?", qty))
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("update fail id not found")
	}

	return nil
}

func (r productRepository) Delete(pId *int) error {
	res := r.db.Model(&model_gorm.Product{}).Delete("id = ?", pId)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("delete fail id not found")
	}

	return nil
}

func (r productRepository) GetProducts() (*[]model_io.Product, error) {
	products := []model_io.Product{}
	err := r.db.Model(&model_gorm.Product{}).Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (r productRepository) GetProductById(pId *int) (*model_io.Product, error) {
	product := model_io.Product{}
	err := r.db.Model(&model_gorm.Product{}).Preload("Category").First(&product, pId).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r productRepository) GetProductsByQuery(q *model_io.QueryProducts) (*[]model_io.Product, error) {

	products := []model_io.Product{}
	skip := (q.Page - 1) * q.Offset
	qName, qCat := utils.CustomQuery(q.Name, q.CategoryId)

	err := r.db.
		Model(&model_gorm.Product{}).
		Where(qName[0], qName[1]).
		Where(qCat[0], qCat[1]).
		Order(q.Sort).
		Limit(q.Offset).
		Offset(skip).
		Preload("Category").
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (r productRepository) GetCounts(q *model_io.QueryProducts) (*int64, error) {
	var count int64
	qName, qCat := utils.CustomQuery(q.Name, q.CategoryId)

	err := r.db.
		Model(&model_gorm.Product{}).
		Where(qName[0], qName[1]).
		Where(qCat[0], qCat[1]).
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &count, nil
}
