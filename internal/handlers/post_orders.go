package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostOrders is a gin handler for the incoming POST /orders request.
func (s *Server) PostOrders(c *gin.Context) {
	orderRequest := &CreateOrderRequest{}
	if err := c.ShouldBindJSON(orderRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewInvalidRequestError("invalid request body"))
		return
	}

	fmt.Printf("Request received: %#v\n", orderRequest)

	err := validateCreateOrderRequest(orderRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	packsSizeCount, err := s.app.CreateOrder(c.Request.Context(), orderRequest.ItemsCount)
	if err != nil {
		log.Default().Printf("app.CreateOrder failed unexpectedly: %v \n", err)
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := mapResponse(packsSizeCount)

	fmt.Println("Request handled successfully")

	c.JSON(http.StatusCreated, response)
}

func validateCreateOrderRequest(orderRequest *CreateOrderRequest) error {
	if orderRequest.ItemsCount <= 0 {
		return NewInvalidRequestError("field 'items_count' cannot be less than 1")
	}

	return nil
}

func mapResponse(packsSizeCount map[int]int) *OrderResponse {
	response := &OrderResponse{
		Packs: make([]*ItemPack, 0),
	}
	for packSize, count := range packsSizeCount {
		response.Packs = append(response.Packs, &ItemPack{Items: packSize, Packs: count})
	}

	return response
}
