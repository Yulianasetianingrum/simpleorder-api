package usecase

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/repository"
)

type CustomerUsecase interface {
	Create(customer *domain.Customer) error
	FindAll(page, limit int, search string) ([]domain.Customer, int64, error)
	FindByID(id uint) (*domain.Customer, error)
	Update(id uint, customer *domain.Customer) error
	Delete(id uint) error
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo}
}

func (u *customerUsecase) Create(customer *domain.Customer) error {
	return u.repo.Create(customer)
}

func (u *customerUsecase) FindAll(page, limit int, search string) ([]domain.Customer, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return u.repo.FindAll(page, limit, search)
}

func (u *customerUsecase) FindByID(id uint) (*domain.Customer, error) {
	return u.repo.FindByID(id)
}

func (u *customerUsecase) Update(id uint, customer *domain.Customer) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	existing.Name = customer.Name
	existing.Email = customer.Email
	existing.Phone = customer.Phone
	existing.Address = customer.Address

	return u.repo.Update(existing)
}

func (u *customerUsecase) Delete(id uint) error {
	_, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	return u.repo.Delete(id)
}
