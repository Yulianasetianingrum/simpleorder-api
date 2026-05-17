package usecase

import (
	"fmt"
	"simpleorder/internal/domain"
	"simpleorder/internal/repository"
	"time"
)

type OrderUsecase interface {
	CreateOrder(req *domain.CreateOrderRequest) (*domain.Order, error)
	FindAll(page, limit int) ([]domain.Order, int64, error)
	FindByID(id uint) (*domain.Order, error)
}

type orderUsecase struct {
	orderRepo    repository.OrderRepository
	customerRepo repository.CustomerRepository
}

func NewOrderUsecase(or repository.OrderRepository, cr repository.CustomerRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo:    or,
		customerRepo: cr,
	}
}

func (u *orderUsecase) CreateOrder(req *domain.CreateOrderRequest) (*domain.Order, error) {
	// Check if customer exists
	_, err := u.customerRepo.FindByID(req.CustomerID)
	if err != nil {
		return nil, domain.ErrCustomerNotFound
	}

	// Generate simple Invoice Number
	invoiceNo := fmt.Sprintf("INV-%d", time.Now().Unix())

	order := &domain.Order{
		InvoiceNumber: invoiceNo,
		CustomerID:    req.CustomerID,
		Status:        "completed", // Auto complete for simplicity
	}

	for _, itemReq := range req.Items {
		order.Items = append(order.Items, domain.OrderItem{
			ProductID: itemReq.ProductID,
			Quantity:  itemReq.Quantity,
		})
	}

	err = u.orderRepo.CreateWithTx(order)
	if err != nil {
		return nil, err
	}

	// Return the saved order with preloaded items if possible, or just the order
	return order, nil
}

func (u *orderUsecase) FindAll(page, limit int) ([]domain.Order, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return u.orderRepo.FindAll(page, limit)
}

func (u *orderUsecase) FindByID(id uint) (*domain.Order, error) {
	return u.orderRepo.FindByID(id)
}
