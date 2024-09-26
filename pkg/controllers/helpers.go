// C:\GoProject\src\ComputerClub\pkg\controllers\errors.go
package controllers

import (
	"ComputerClub/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// handleError обрабатывает все ошибки, возникающие в процессе выполнения...
// Добавляет статус код к ним и сразу возвращает клиенту...
func handleError(c *gin.Context, err error) {
	switch {
	// Ошибки аутентификации
	case errors.Is(err, errs.ErrEmptyAuthHeader):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrInvalidAuthHeader):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrTokenParsingFailed):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotAuthenticated):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// Ошибки разрешений
	case errors.Is(err, errs.ErrPermissionDenied):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrPermissionDeniedOnlyForAdmin):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrPermissionDeniedOnlyForSuperAdmin):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	// Ошибки пользователей
	case errors.Is(err, errs.ErrUserNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUsersNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUsernameUniquenessFailed):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserAlreadyDeleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotDeleted):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrSomethingWentWrong):
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))

	// Ошибки категорий
	case errors.Is(err, errs.ErrCategoryNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCategoryNameExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Ошибки компьютеров
	case errors.Is(err, errs.ErrComputerNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrComputerAlreadyExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Ошибки бронирования
	case errors.Is(err, errs.ErrBookingNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrBookingAlreadyCompleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Ошибки баланса
	case errors.Is(err, errs.ErrBalanceInsufficient):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrBalanceNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	// Общие ошибки
	case errors.Is(err, errs.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrServerError):
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))

	// Внутренняя ошибка сервера
	default:
		c.JSON(http.StatusInternalServerError, newErrorResponse("internal server error"))
	}
}

// ErrorResponse представляет структуру для обработки сообщений об ошибках
type ErrorResponse struct {
	Error string `json:"error"` // Описание возникшей ошибки...
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
