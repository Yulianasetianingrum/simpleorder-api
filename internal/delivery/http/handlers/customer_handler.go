package handlers

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/usecase"
	"simpleorder/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	usecase  usecase.CustomerUsecase
	validate *validator.Validate
}

func NewCustomerHandler(u usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		usecase:  u,
		validate: validator.New(),
	}
}

// Create godoc
// @Summary Create a new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param request body domain.Customer true "Customer Data"
// @Security BearerAuth
// @Success 201 {object} response.Response
// @Router /api/v1/customers [post]
func (h *CustomerHandler) Create(c *fiber.Ctx) error {
	var customer domain.Customer
	if err := c.BodyParser(&customer); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validate.Struct(customer); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	if err := h.usecase.Create(&customer); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create customer", err)
	}

	return response.Success(c, fiber.StatusCreated, "Customer created successfully", customer)
}

// FindAll godoc
// @Summary Get all customers
// @Tags Customers
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name or email"
// @Security BearerAuth
// @Success 200 {object} response.PaginatedResponse
// @Router /api/v1/customers [get]
func (h *CustomerHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	customers, total, err := h.usecase.FindAll(page, limit, search)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch customers", err)
	}

	return response.SuccessPaginated(c, fiber.StatusOK, "Customers fetched successfully", customers, total, page, limit)
}
