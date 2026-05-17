package handlers

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/usecase"
	"simpleorder/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase  usecase.ProductUsecase
	validate *validator.Validate
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase:  u,
		validate: validator.New(),
	}
}

// Create godoc
// @Summary Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body domain.Product true "Product Data"
// @Security BearerAuth
// @Success 201 {object} response.Response
// @Router /api/v1/products [post]
func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validate.Struct(product); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	if err := h.usecase.Create(&product); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create product", err)
	}

	return response.Success(c, fiber.StatusCreated, "Product created successfully", product)
}

// FindAll godoc
// @Summary Get all products
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name"
// @Security BearerAuth
// @Success 200 {object} response.PaginatedResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	products, total, err := h.usecase.FindAll(page, limit, search)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch products", err)
	}

	return response.SuccessPaginated(c, fiber.StatusOK, "Products fetched successfully", products, total, page, limit)
}
