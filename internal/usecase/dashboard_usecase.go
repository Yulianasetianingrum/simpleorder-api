package usecase

import "simpleorder/internal/repository"

type DashboardStats struct {
	TotalOrders    int64   `json:"total_orders"`
	TotalRevenue   float64 `json:"total_revenue"`
	TotalCustomers int64   `json:"total_customers"`
}

type DashboardUsecase interface {
	GetStats() (*DashboardStats, error)
}

type dashboardUsecase struct {
	orderRepo repository.OrderRepository
}

func NewDashboardUsecase(or repository.OrderRepository) DashboardUsecase {
	return &dashboardUsecase{orderRepo: or}
}

func (u *dashboardUsecase) GetStats() (*DashboardStats, error) {
	orders, revenue, customers, err := u.orderRepo.GetStats()
	if err != nil {
		return nil, err
	}

	return &DashboardStats{
		TotalOrders:    orders,
		TotalRevenue:   revenue,
		TotalCustomers: customers,
	}, nil
}
