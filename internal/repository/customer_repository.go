package repository

import (
	"simpleorder/internal/domain"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *domain.Customer) error
	FindAll(page, limit int, search string) ([]domain.Customer, int64, error)
	FindByID(id uint) (*domain.Customer, error)
	Update(customer *domain.Customer) error
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) Create(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) FindAll(page, limit int, search string) ([]domain.Customer, int64, error) {
	var customers []domain.Customer
	var total int64

	query := r.db.Model(&domain.Customer{})

	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at desc").Find(&customers).Error

	return customers, total, err
}

func (r *customerRepository) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrCustomerNotFound
		}
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}
