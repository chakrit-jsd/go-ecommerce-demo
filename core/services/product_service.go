package services

import (
	"fmt"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"
)

type productService struct {
	prodRepo ports.ProductRepository
}

func NewProductService(prodRepo ports.ProductRepository) ports.ProductService {
	return productService{prodRepo}
}

func (s productService) Create(pd *model_io.Product) error {

	product := model_gorm.Product{
		Name:       pd.Name,
		Detail:     pd.Detail,
		Stock:      pd.Stock,
		Price:      pd.Price,
		CategoryId: pd.CategoryId,
	}

	err := s.prodRepo.Create(&product)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s productService) Update(pd *model_io.Product, pId *int) error {
	product := model_gorm.Product{
		Name:       pd.Name,
		Detail:     pd.Detail,
		Stock:      pd.Stock,
		Price:      pd.Price,
		CategoryId: pd.CategoryId,
	}

	err := s.prodRepo.Update(&product, pId)
	if err != nil {
		return utils.CusErrorDB(err)
	}
	return nil
}

func (s productService) UpdateStock(qty, pId *int) error {
	err := s.prodRepo.UpdateStock(qty, pId)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s productService) Delete(pId *int) error {
	err := s.prodRepo.Delete(pId)
	if err != nil {
		return utils.CusErrorDB(err)
	}
	return nil
}

func (s productService) GetProducts() (*[]model_io.Product, error) {
	products, err := s.prodRepo.GetProducts()
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return products, nil
}

func (s productService) GetProductById(pId *int) (*model_io.Product, error) {
	product, err := s.prodRepo.GetProductById(pId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return product, nil
}

func (s productService) GetProductsByQuery(q *model_io.QueryProducts, count bool) (*model_io.ProductsAndCounts, error) {
	ProdAndCounts := model_io.ProductsAndCounts{
		TotalCounts: -1,
	}
	if q.Sort == "" {
		q.Sort = "price"
	}
	if q.Offset == 0 {
		q.Offset = 10
	}

	len := len(q.Sort)
	if q.Sort[len-4:] == "desc" {
		q.Sort = fmt.Sprintf("%v %v", q.Sort[:len-4], q.Sort[len-4:])
	}

	if count {
		c, err := s.prodRepo.GetCounts(q)
		if err != nil {
			return nil, utils.CusErrorDB(err)
		}
		ProdAndCounts.TotalCounts = *c
	}

	products, err := s.prodRepo.GetProductsByQuery(q)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	ProdAndCounts.Products = *products

	return &ProdAndCounts, nil
}
