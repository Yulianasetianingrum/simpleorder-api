package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	TotalPages int         `json:"total_pages"`
}

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessPaginated(c *fiber.Ctx, statusCode int, message string, data interface{}, total int64, page, limit int) error {
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return c.Status(statusCode).JSON(PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Total:      total,
		Page:       page,
		TotalPages: totalPages,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string, err error) error {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
		Error:   errStr,
	})
}
