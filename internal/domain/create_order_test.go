package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/domain/mocks"
)

func TestCreateOrderSuccess(t *testing.T) {
	// Arrange
	testCases := []struct {
		availablePacks   []int
		itemsOrdered     int
		expectedResponse map[int]int
	}{
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     1,
			expectedResponse: map[int]int{250: 1},
		},
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     250,
			expectedResponse: map[int]int{250: 1},
		},
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     251,
			expectedResponse: map[int]int{500: 1},
		},
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     501,
			expectedResponse: map[int]int{500: 1, 250: 1},
		},
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     12001,
			expectedResponse: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			availablePacks:   []int{250, 500, 1000, 2000, 5000},
			itemsOrdered:     12251,
			expectedResponse: map[int]int{5000: 2, 2000: 1, 500: 1},
		},
		{
			availablePacks:   []int{23, 31, 53},
			itemsOrdered:     500000,
			expectedResponse: map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("ordered_%d_items_available_packs_%v", testCase.itemsOrdered, testCase.availablePacks)
		t.Run(testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := mocks.NewMockRepository(ctrl)
			app, err := NewApplication(testCase.availablePacks, mockRepo)
			mockRepo.EXPECT().StoreOrder(gomock.Any(), gomock.Any()).Times(1)
			assert.NoError(t, err)
			itemsOrdered := testCase.itemsOrdered

			// Act
			response, err := app.CreateOrder(t.Context(), itemsOrdered)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedResponse, response)
		})
	}
}

func TestCreateOrderWhenRepositoryFailsThenError(t *testing.T) {
	// Arrange

	availablePacks := []int{250, 500, 1000, 2000, 5000}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	app, err := NewApplication(availablePacks, mockRepo)
	assert.NoError(t, err)
	mockErr := fmt.Errorf("repository error")
	mockRepo.EXPECT().StoreOrder(gomock.Any(), gomock.Any()).Times(1).Return(mockErr)

	// Act
	response, err := app.CreateOrder(t.Context(), 100)

	// Assert
	assert.Nil(t, response)
	assert.Equal(t, mockErr, err)
}
