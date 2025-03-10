package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/domain"
	"repartnerstask.com/m/internal/handlers"
	"repartnerstask.com/m/internal/handlers/mocks"
)

type ServerTestSuite struct {
	suite.Suite

	mockApp *mocks.MockApplication
	server  *handlers.Server
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (s *ServerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.mockApp = mocks.NewMockApplication(ctrl)
	s.server = handlers.NewServer(s.mockApp)
}

func (s *ServerTestSuite) TestPostOrdersSuccess() {
	// Arrange
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.POST("/orders", s.server.PostOrders)

	requestBody := strings.NewReader(`{"items_count":750}`)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/orders", requestBody)

	mockResult := map[int]int{500: 1, 250: 1}
	s.mockApp.EXPECT().CreateOrder(gomock.Any(), 750).Times(1).Return(mockResult, nil)

	expectedResponse := &handlers.OrderResponse{
		Packs: []*handlers.ItemPack{
			{
				Items: 500,
				Packs: 1,
			},
			{
				Items: 250,
				Packs: 1,
			},
		},
	}

	// Act
	engine.ServeHTTP(w, ctx.Request)
	defer w.Result().Body.Close()

	// Assert
	s.Equal(http.StatusCreated, w.Result().StatusCode)
	responseBody := &handlers.OrderResponse{}
	err := json.NewDecoder(w.Result().Body).Decode(responseBody)
	s.NoError(err)

	for _, expectedPackCount := range expectedResponse.Packs {
		found := false

		for _, actualPackCount := range responseBody.Packs {
			if expectedPackCount.Items == actualPackCount.Items &&
				expectedPackCount.Packs == actualPackCount.Packs {
				found = true
				break
			}
		}

		s.True(found, fmt.Sprintf("expected items pack not found: %d \n", expectedPackCount.Items))
	}
}

func (s *ServerTestSuite) TestPostOrdersWhenInvalidRequestThenReturn400Error() {
	// Arrange
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.POST("/orders", s.server.PostOrders)
	msgInvalidRequestBody := "{\"message\":\"invalid request body\"}"

	testCases := []struct {
		name             string
		requestBody      io.Reader
		expectedResponse string
	}{
		{
			name:             "body is nil",
			requestBody:      nil,
			expectedResponse: msgInvalidRequestBody,
		},
		{
			name:             "body is empty string",
			requestBody:      strings.NewReader(""),
			expectedResponse: msgInvalidRequestBody,
		},
		{
			name:             "body is invalid json",
			requestBody:      strings.NewReader(`{"items_count":0`),
			expectedResponse: msgInvalidRequestBody,
		},
		{
			name:             "items_count is 0",
			requestBody:      strings.NewReader(`{"items_count":0}`),
			expectedResponse: "{\"message\":\"field 'items_count' cannot be less than 1\"}",
		},
		{
			name:             "items_count is -1",
			requestBody:      strings.NewReader(`{"items_count":-1}`),
			expectedResponse: "{\"message\":\"field 'items_count' cannot be less than 1\"}",
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			ctx.Request = httptest.NewRequest(http.MethodPost, "/orders", testCase.requestBody)
			w := httptest.NewRecorder()

			// Act
			engine.ServeHTTP(w, ctx.Request)
			defer w.Result().Body.Close()

			// Assert
			s.Equal(http.StatusBadRequest, w.Result().StatusCode)
			responseBodyBytes, err := io.ReadAll(w.Result().Body)
			s.NoError(err)
			s.Equal(testCase.expectedResponse, string(responseBodyBytes))
		})
	}
}

func (s *ServerTestSuite) TestPostOrdersWhenApplicationFailsThenReturn500Error() {
	// Arrange
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.POST("/orders", s.server.PostOrders)
	requestBody := strings.NewReader(`{"items_count":250}`)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/orders", requestBody)

	s.mockApp.EXPECT().CreateOrder(gomock.Any(), 250).Times(1).Return(nil, &domain.CannotCalculateOrderItemsError{})

	// Act
	engine.ServeHTTP(w, ctx.Request)
	defer w.Result().Body.Close()

	// Assert
	s.Equal(http.StatusInternalServerError, w.Result().StatusCode)
	responseBodyBytes, err := io.ReadAll(w.Result().Body)
	s.NoError(err)
	s.Empty(string(responseBodyBytes))
}
