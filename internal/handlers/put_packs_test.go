package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"repartnerstask.com/m/internal/handlers"
	"repartnerstask.com/m/internal/handlers/mocks"
)

func TestPutPacks(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockApp := mocks.NewMockApplication(ctrl)
	server := handlers.NewServer(mockApp)
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.PUT("/packs", server.PutPacks)

	requestBody := strings.NewReader(`{"packs":[23, 31, 53]}`)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/packs", requestBody)

	mockApp.EXPECT().ReplacePacks([]int{23, 31, 53}).Times(1)

	// Act
	engine.ServeHTTP(w, ctx.Request)
	defer w.Result().Body.Close()

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
}
