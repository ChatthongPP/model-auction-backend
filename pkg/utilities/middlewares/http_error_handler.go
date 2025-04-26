package middlewares

import (
	"backend-service/pkg/utilities/responses"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(logger *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var appErr *responses.ApplicationError

		if errors.As(err, &appErr) {
			logger.WithError(err).WithField("code", appErr.Code).Error(err.Error())
			if err := c.JSON(responses.GetHttpStatusForCode(appErr.Code), responses.ErrorResponse{
				Code: appErr.Code,
				Error: responses.ErrorDetail{
					Message: appErr.Message,
					Stack:   appErr.Error(),
				},
			}); err != nil {
				logger.Error(err)
			}
		}

		logger.WithError(err).Error(err.Error())

		if err := c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Code: responses.UNEXPECTED_EXCEPTION,
			Error: responses.ErrorDetail{
				Message: responses.DEFAULT_ERROR_MESSAGE,
				Stack:   err.Error(),
			},
		}); err != nil {
			logger.Error(err)
		}
	}
}
