package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/domain/mocks"
	"repartnerstask.com/m/internal/domain/types"
)

func TestGetOrdersSuccess(t *testing.T) {
	// Arrange
	availablePacks := []int{250, 500, 1000, 2000, 5000}
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	app, err := NewApplication(availablePacks, mockRepo)
	ordersCount := 20
	mockOrders := []*types.Order{
		{
			ItemsShipped: 10,
			PacksShipped: 1,
			ItemPacks:    map[int]int{10: 1},
		},
	}
	mockRepo.EXPECT().ListOrdersDesc(gomock.Any(), ordersCount).Times(1).Return(mockOrders, nil)
	assert.NoError(t, err)

	// Act
	response, err := app.GetOrders(t.Context(), ordersCount)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, mockOrders, response)
}

func TestGetOrdersWhenRepositoryFailsThenError(t *testing.T) {
	// Arrange
	availablePacks := []int{250, 500, 1000, 2000, 5000}
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	app, err := NewApplication(availablePacks, mockRepo)
	ordersCount := 20
	mockErr := fmt.Errorf("repository error")
	mockRepo.EXPECT().ListOrdersDesc(gomock.Any(), ordersCount).Times(1).Return(nil, mockErr)
	assert.NoError(t, err)

	// Act
	response, err := app.GetOrders(t.Context(), ordersCount)

	// Assert
	assert.Equal(t, mockErr, err)
	assert.Nil(t, response)
}
