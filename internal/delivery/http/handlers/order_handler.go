package handlers

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/usecase"
	"simpleorder/pkg/response"
	"simpleorder/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	usecase  usecase.OrderUsecase
	validate *validator.Validate
}

func NewOrderHandler(u usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		usecase:  u,
		validate: validator.New(),
	}
}

// Create godoc
// @Summary Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body domain.CreateOrderRequest true "Order Data"
// @Security BearerAuth
// @Success 201 {object} response.Response
// @Router /api/v1/orders [post]
func (h *OrderHandler) Create(c *fiber.Ctx) error {
	var req domain.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	order, err := h.usecase.CreateOrder(&req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create order", err)
	}

	return response.Success(c, fiber.StatusCreated, "Order created successfully", order)
}

// FindAll godoc
// @Summary Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Security BearerAuth
// @Success 200 {object} response.PaginatedResponse
// @Router /api/v1/orders [get]
func (h *OrderHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	orders, total, err := h.usecase.FindAll(page, limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch orders", err)
	}

	return response.SuccessPaginated(c, fiber.StatusOK, "Orders fetched successfully", orders, total, page, limit)
}

// GenerateInvoice godoc
// @Summary Generate Order Invoice PDF
// @Tags Orders
// @Produce application/pdf
// @Param id path int true "Order ID"
// @Security BearerAuth
// @Success 200 {file} file
// @Router /api/v1/orders/{id}/invoice [get]
func (h *OrderHandler) GenerateInvoice(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid order ID", err)
	}

	order, err := h.usecase.FindByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Order not found", err)
	}

	pdf, err := utils.GenerateInvoicePDF(order)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate PDF", err)
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", `attachment; filename="invoice_`+order.InvoiceNumber+`.pdf"`)

	return pdf.Output(c.Response().BodyWriter())
}
