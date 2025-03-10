package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutPacks handles requests to replace pack size configuration for the application.
// Returns a 400 status code if any of the pack sizes is less than or equal to 0.
func (s *Server) PutPacks(c *gin.Context) {
	request := &ReplacePacksRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewInvalidRequestError("invalid request body"))
		return
	}

	s.app.ReplacePacks(request.Packs)

	c.Status(http.StatusNoContent)
}
