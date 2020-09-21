package usecases

import (
	"myfirstgosql/models"
	"myfirstgosql/repositories"
)

type IProductUseCase interface {
	GetProductById(id string) (*models.Product, error)
	GetProductByNameLike(name []string) ([]*models.Product, error)
	GetProductPaging(pageNo, totalPerPage int) ([]*models.Product, error)
	GetTotalProduct() (int64, error)
	RegisterNewProduct(product models.Product) (*models.Product, error)
}

type ProductUseCase struct {
	repo repositories.IProductRepository
}

func (p *ProductUseCase) RegisterNewProduct(product models.Product) (*models.Product, error) {
	return p.repo.Insert(product)
}

func NewProductUseCase(repo repositories.IProductRepository) IProductUseCase {
	return &ProductUseCase{repo}
}

func (p *ProductUseCase) GetProductById(id string) (*models.Product, error) {
	return p.repo.FindOneById(id)
}

func (p *ProductUseCase) GetProductByNameLike(name []string) ([]*models.Product, error) {
	result := make([]*models.Product, 0)
	for _, q := range name {
		r, err := p.repo.FindAllByNameLike(q)
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}
	return result, nil
}

func (p *ProductUseCase) GetProductPaging(pageNo, totalPerPage int) ([]*models.Product, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllProductPaging((pageNo-1)*totalPerPage, totalPerPage)
}
func (p *ProductUseCase) GetTotalProduct() (int64, error) {
	return p.repo.Count()
}
