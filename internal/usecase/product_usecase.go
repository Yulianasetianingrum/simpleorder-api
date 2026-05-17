package usecase

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/repository"
)

type ProductUsecase interface {
	Create(product *domain.Product) error
	FindAll(page, limit int, search string) ([]domain.Product, int64, error)
	FindByID(id uint) (*domain.Product, error)
	Update(id uint, product *domain.Product) error
	Delete(id uint) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo}
}

func (u *productUsecase) Create(product *domain.Product) error {
	return u.repo.Create(product)
}

func (u *productUsecase) FindAll(page, limit int, search string) ([]domain.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return u.repo.FindAll(page, limit, search)
}

func (u *productUsecase) FindByID(id uint) (*domain.Product, error) {
	return u.repo.FindByID(id)
}

func (u *productUsecase) Update(id uint, product *domain.Product) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	existing.Name = product.Name
	existing.Description = product.Description
	existing.Price = product.Price
	existing.Stock = product.Stock

	return u.repo.Update(existing)
}

func (u *productUsecase) Delete(id uint) error {
	_, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	return u.repo.Delete(id)
}
