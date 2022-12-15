package v1

import (
	"context"
	"net/http"

	"github.com/TemurMannonov/medium_api_gateway/api/models"

	pbp "github.com/TemurMannonov/medium_api_gateway/genproto/post_service"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		req models.CreateCategoryRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.CategoryService().Create(context.Background(), &pbp.Category{
		Title: req.Title,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.Category{
		ID:        resp.Id,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}
