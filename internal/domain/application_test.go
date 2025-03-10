package domain

// intentionally not using the <filename>_test.go to do "white box" testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/domain/mocks"
)

func TestNewApplicationSuccess(t *testing.T) {
	// Arrange
	packs := []int{250, 500, 1000, 2000, 5000}
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)

	// Act
	app, err := NewApplication(packs, mockRepo)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, packs, app.packSizes)
}

func TestNewApplicationWhenInvalidPacksThenError(t *testing.T) {
	// Arrange
	packs := []int{}
	expectedError := &NoPacksError{}
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)

	// Act
	app, err := NewApplication(packs, mockRepo)

	// Assert
	assert.Empty(t, app)
	assert.Equal(t, expectedError, err)
}
