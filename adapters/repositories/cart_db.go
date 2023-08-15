package repositories

import (
	"errors"
	"fmt"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"

	"gorm.io/gorm"
)

type cartRepositoryDB struct {
	db *gorm.DB
}

func NewCartRepositoryDB(db *gorm.DB) ports.CartRepository {
	return cartRepositoryDB{db}
}

func (r cartRepositoryDB) GetCart(userId *int) (*[]model_io.CartDetail, error) {
	cartDetails := []model_io.CartDetail{}
	err := r.db.Model(&model_gorm.CartDetail{}).
		Where("user_id = ?", *userId).
		Preload("Product").
		Preload("Product.Category").
		Find(&cartDetails).Error
	if err != nil {
		return nil, err
	}

	return &cartDetails, nil
}

func (r cartRepositoryDB) CreateCart(userId *int) error {
	return nil
}

func (r cartRepositoryDB) AddProduct(cart *model_gorm.CartDetail) error {

	res := r.db.Model(&model_gorm.CartDetail{}).
		Where("user_id = ? AND product_id = ?", cart.UserId, cart.ProductId).
		Update("quantity", gorm.Expr("quantity + ?", 1))
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		cart.Quantity = 1
		err := r.db.Model(&model_gorm.CartDetail{}).Create(&cart).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (r cartRepositoryDB) UpdateProduct(cart *model_gorm.CartDetail) error {

	res := r.db.Model(&model_gorm.CartDetail{}).
		Where("user_id = ? AND product_id = ?", cart.UserId, cart.ProductId).
		Update("quantity", cart.Quantity)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (r cartRepositoryDB) DeleteProduct(cart *model_io.DeleteProductsInCart) error {
	res := r.db.Model(&model_gorm.CartDetail{}).
		Where("user_id = ? AND product_id IN ?", cart.UserId, cart.ProductsId).
		Delete("")
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
