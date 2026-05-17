package handlers

import (
	"simpleorder/internal/config"
	"simpleorder/internal/domain"
	"simpleorder/internal/usecase"
	"simpleorder/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	cfg         *config.Config
	validate    *validator.Validate
}

func NewAuthHandler(au usecase.AuthUsecase, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authUsecase: au,
		cfg:         cfg,
		validate:    validator.New(),
	}
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.UserRegistrationRequest true "Register User"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.UserRegistrationRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	err := h.authUsecase.Register(&req)
	if err != nil {
		if err == domain.ErrUsernameExists {
			return response.Error(c, fiber.StatusConflict, "Username already exists", err)
		}
		return response.Error(c, fiber.StatusInternalServerError, "Failed to register user", err)
	}

	return response.Success(c, fiber.StatusCreated, "User registered successfully", nil)
}

// Login godoc
// @Summary Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.UserLoginRequest true "Login User"
// @Success 200 {object} response.Response{data=domain.UserLoginResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	res, err := h.authUsecase.Login(&req, h.cfg.JWTSecret, h.cfg.JWTExpHours)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid credentials", err)
	}

	return response.Success(c, fiber.StatusOK, "Login successful", res)
}
