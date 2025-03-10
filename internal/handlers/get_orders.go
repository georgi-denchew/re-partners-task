package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"repartnerstask.com/m/internal/domain/types"
)

// GetOrders is a handler that returns orders. Takes in an optional 'count' param, which defaults to 20.
func (s *Server) GetOrders(c *gin.Context) {
	countStr := c.DefaultQuery("count", "20")
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewInvalidRequestError("count query parameter must be a positive integer"))
		return
	}

	fmt.Printf("GetOrders request received, count= %d\n", count)

	orders, err := s.app.GetOrders(c.Request.Context(), count)
	if err != nil {
		log.Default().Printf("app.GetOrders failed unexpectedly: %v \n", err)
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := mapGetResponse(orders)

	fmt.Println("Request handled successfully")

	c.JSON(http.StatusOK, response)
}

func mapGetResponse(orders []*types.Order) []*OrderResponse {
	response := make([]*OrderResponse, 0)
	for _, order := range orders {
		mappedOrder := &OrderResponse{
			CreatedAt: &order.CreatedAt,
			Packs:     make([]*ItemPack, 0),
		}

		for packSize, packCount := range order.ItemPacks {
			mappedOrder.Packs = append(mappedOrder.Packs, &ItemPack{Items: packSize, Packs: packCount})
		}

		response = append(response, mappedOrder)
	}

	return response
}
