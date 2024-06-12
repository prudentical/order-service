package api

import (
	"log/slog"
	"net/http"
	"order-service/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PositionHandler interface {
	GetByBotId(c echo.Context) error
	DeleteByBotId(c echo.Context) error
	HandleRoutes(e *echo.Echo)
}

type PositionHandlerImpl struct {
	service service.PositionService
	logger  *slog.Logger
}

func NewPositionHandler(service service.PositionService, logger *slog.Logger) PositionHandler {
	return PositionHandlerImpl{service, logger}
}

func (h PositionHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/users/:user_id/accounts/:account_id/bots/:bot_id/positions", h.GetByBotId)
	e.DELETE("/users/:user_id/accounts/:account_id/bots/:bot_id/positions", h.DeleteByBotId)
}

func (h PositionHandlerImpl) GetByBotId(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: userIdStr, TypeName: "User"}
	}

	accountIdStr := c.Param("account_id")
	accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: accountIdStr, TypeName: "Bot"}
	}

	botIdStr := c.Param("bot_id")
	botId, err := strconv.ParseInt(botIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: botIdStr, TypeName: "Bot"}
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
	status := c.QueryParam("status")

	result, err := h.service.GetByBotId(userId, accountId, botId, page, size, status)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (h PositionHandlerImpl) DeleteByBotId(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: userIdStr, TypeName: "User"}
	}

	accountIdStr := c.Param("account_id")
	accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: accountIdStr, TypeName: "Bot"}
	}

	botIdStr := c.Param("bot_id")
	botId, err := strconv.ParseInt(botIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{Id: botIdStr, TypeName: "Bot"}
	}
	err = h.service.DeleteByBotId(userId, accountId, botId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
