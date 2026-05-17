package handlers

import (
	"simpleorder/internal/usecase"
	"simpleorder/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	usecase usecase.DashboardUsecase
}

func NewDashboardHandler(u usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{usecase: u}
}

// GetStats godoc
// @Summary Get Dashboard Stats
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=usecase.DashboardStats}
// @Router /api/v1/dashboard/stats [get]
func (h *DashboardHandler) GetStats(c *fiber.Ctx) error {
	stats, err := h.usecase.GetStats()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch stats", err)
	}

	return response.Success(c, fiber.StatusOK, "Dashboard stats fetched successfully", stats)
}
