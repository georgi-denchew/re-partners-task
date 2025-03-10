package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/domain/types"
	"repartnerstask.com/m/internal/handlers"
	"repartnerstask.com/m/internal/handlers/mocks"
)

func TestGetOrdersSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockApp := mocks.NewMockApplication(ctrl)
	server := handlers.NewServer(mockApp)
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.GET("/orders", server.GetOrders)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/orders?count=5", nil)
	createdAt := time.Now().UTC()
	createdAt2 := time.Now().UTC().Add(time.Hour * -1)

	mockResponse := []*types.Order{
		{
			ItemsShipped: 100,
			PacksShipped: 1,
			CreatedAt:    createdAt,
			ItemPacks:    map[int]int{100: 1},
		},
		{
			ItemsShipped: 200,
			PacksShipped: 2,
			CreatedAt:    createdAt2,
			ItemPacks:    map[int]int{200: 1},
		},
	}
	mockApp.EXPECT().GetOrders(gomock.Any(), 5).Return(mockResponse, nil)

	expectedResponse := []*handlers.OrderResponse{
		{
			CreatedAt: &createdAt,
			Packs: []*handlers.ItemPack{
				{
					Items: 100,
					Packs: 1,
				},
			},
		},
		{
			CreatedAt: &createdAt2,
			Packs: []*handlers.ItemPack{
				{
					Items: 200,
					Packs: 1,
				},
			},
		},
	}

	// Act
	engine.ServeHTTP(w, ctx.Request)
	defer w.Result().Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)

	var responseBody []handlers.OrderResponse
	err := json.NewDecoder(w.Result().Body).Decode(&responseBody)

	assert.NoError(t, err)

	for _, expectedOrder := range expectedResponse {
		found := false
		for _, actualOrder := range responseBody {
			if *expectedOrder.CreatedAt == *actualOrder.CreatedAt {
				found = true
				for i := range expectedOrder.Packs {
					assert.Equal(t, expectedOrder.Packs[i].Items, actualOrder.Packs[i].Items)
					assert.Equal(t, expectedOrder.Packs[i].Packs, actualOrder.Packs[i].Packs)
				}
				break
			}
		}

		assert.True(t, found, "expected items pack not found:")
	}
}

func TestGetOrdersWhenApplicationFailsThenError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockApp := mocks.NewMockApplication(ctrl)
	server := handlers.NewServer(mockApp)
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.GET("/orders", server.GetOrders)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/orders?count=5", nil)

	mockErr := fmt.Errorf("application error")
	mockApp.EXPECT().GetOrders(gomock.Any(), 5).Return(nil, mockErr)

	// Act
	engine.ServeHTTP(w, ctx.Request)
	defer w.Result().Body.Close()

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
