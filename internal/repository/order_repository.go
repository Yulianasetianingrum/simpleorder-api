package repository

import (
	"simpleorder/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateWithTx(order *domain.Order) error
	FindAll(page, limit int) ([]domain.Order, int64, error)
	FindByID(id uint) (*domain.Order, error)
	GetStats() (int64, float64, int64, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateWithTx(order *domain.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Calculate total amount and check stock
		var totalAmount float64

		for i, item := range order.Items {
			var product domain.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return domain.ErrProductNotFound
			}

			if product.Stock < item.Quantity {
				return domain.ErrInsufficientStock
			}

			// Reduce stock
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			order.Items[i].Price = product.Price
			order.Items[i].SubTotal = product.Price * float64(item.Quantity)
			totalAmount += order.Items[i].SubTotal
		}

		order.TotalAmount = totalAmount

		// Create Order and Items
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *orderRepository) FindAll(page, limit int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	var total int64

	query := r.db.Model(&domain.Order{})
	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("Customer").Offset(offset).Limit(limit).Order("created_at desc").Find(&orders).Error

	return orders, total, err
}

func (r *orderRepository) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Customer").Preload("Items").Preload("Items.Product").First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetStats() (int64, float64, int64, error) {
	var totalOrders int64
	var totalRevenue float64
	var totalCustomers int64

	r.db.Model(&domain.Order{}).Count(&totalOrders)
	
	type Result struct {
		Total float64
	}
	var res Result
	r.db.Model(&domain.Order{}).Select("sum(total_amount) as total").Scan(&res)
	totalRevenue = res.Total

	r.db.Model(&domain.Customer{}).Count(&totalCustomers)

	return totalOrders, totalRevenue, totalCustomers, nil
}
