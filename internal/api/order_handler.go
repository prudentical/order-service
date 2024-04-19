package api

import (
	"log/slog"
	"net/http"
	"order-service/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	GetByBotId(c echo.Context) error
	DeleteByBotId(c echo.Context) error
	HandleRoutes(e *echo.Echo)
}

type OrderHandlerImpl struct {
	service service.OrderService
	logger  *slog.Logger
}

func NewOrderHandler(service service.OrderService, logger *slog.Logger) OrderHandler {
	return OrderHandlerImpl{service, logger}
}

func (h OrderHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/bots/:bot_id/orders", h.GetByBotId)
	e.DELETE("/bots/:bot_id/orders", h.DeleteByBotId)
}

func (h OrderHandlerImpl) GetByBotId(c echo.Context) error {
	botIdStr := c.Param("bot_id")
	botId, err := strconv.Atoi(botIdStr)
	if err != nil {
		return err
	}
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	sizeStr := c.QueryParam("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 20
	}
	result, err := h.service.GetByBotId(botId, page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (h OrderHandlerImpl) DeleteByBotId(c echo.Context) error {
	botIdStr := c.Param("bot_id")
	botId, err := strconv.Atoi(botIdStr)
	if err != nil {
		return err
	}
	err = h.service.DeleteByBotId(botId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
