package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"

	"userService/internal/gen"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

type APIHandler struct {
	Handler *Handler
}

func (h *APIHandler) SaveUser(ctx echo.Context) error {
	var u api.User
	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.Handler.DB.Where("external_id = ?", u.ExternalId).Assign(&u).FirstOrCreate(&u).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save user"})
	}

	return ctx.JSON(http.StatusCreated, u)
}

func (h *APIHandler) GetUser(ctx echo.Context, id openapi_types.UUID) error {
	var user api.User

	if err := h.Handler.DB.First(&user, "external_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}

	return ctx.JSON(http.StatusOK, user)
}
