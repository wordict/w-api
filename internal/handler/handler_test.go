//go:build unit

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wordict/w-api/internal/entity"
	"github.com/wordict/w-api/internal/handler"
	"github.com/wordict/w-api/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createHandlerMocks(t *testing.T) (*mocks.MockLogger, *mocks.MockService) {
	ctrl := gomock.NewController(t)
	return mocks.NewMockLogger(ctrl), mocks.NewMockService(ctrl)
}
func createFiberAppForHandler(h *handler.Handler) *fiber.App {
	app := fiber.New()
	app.Post("/signup", h.Signup)
	return app
}
func makeTestRequestWithBody(method string, route string, body interface{}) *http.Request {
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(method, route, bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	return req
}
func TestHandler_Signup(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	handler := handler.New(mockService, mockLogger)
	app := createFiberAppForHandler(handler)
	t.Run("given invalid signup request then it should return status bad request", func(t *testing.T) {
		givenRequest := "test"
		req := makeTestRequestWithBody(fiber.MethodPost, "/signup", givenRequest)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("given valid signup request and error occurred on service then it should return status internal server error", func(t *testing.T) {
		givenRequest := entity.SignupRequest{
			Email:    "first@gmail.com",
			Password: "****",
		}
		mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any())
		mockService.EXPECT().Signup(gomock.Any(), givenRequest).Return(errors.New("any"))
		req := makeTestRequestWithBody(fiber.MethodPost, "/signup", givenRequest)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("given valid signup request and error occurred on service then it should return status internal server error", func(t *testing.T) {
		givenRequest := entity.SignupRequest{
			Email:    "first@gmail.com",
			Password: "****",
		}
		mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any())
		mockService.EXPECT().Signup(gomock.Any(), givenRequest).Return(nil)
		req := makeTestRequestWithBody(fiber.MethodPost, "/signup", givenRequest)
		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
