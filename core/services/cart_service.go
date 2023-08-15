package services

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"
)

type cartService struct {
	cartRepo ports.CartRepository
}

func NewCartService(cartRepo ports.CartRepository) ports.CartService {
	return cartService{cartRepo}
}

func (s cartService) GetCart(userId *int) (*[]model_io.CartDetail, error) {
	cart, err := s.cartRepo.GetCart(userId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return cart, nil
}

func (s cartService) AddProduct(c *model_io.CartDetail) error {
	cart := model_gorm.CartDetail{
		UserId:    c.UserId,
		ProductId: c.ProductId,
	}
	err := s.cartRepo.AddProduct(&cart)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s cartService) UpdateProduct(c *model_io.CartDetail) error {
	cart := model_gorm.CartDetail{
		UserId:    c.UserId,
		ProductId: c.ProductId,
		Quantity:  c.Quantity,
	}
	err := s.cartRepo.UpdateProduct(&cart)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s cartService) DeleteProduct(cart *model_io.DeleteProductsInCart) error {
	err := s.cartRepo.DeleteProduct(cart)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}
