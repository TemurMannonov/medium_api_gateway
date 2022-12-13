package v1

import (
	"context"
	"net/http"

	"github.com/TemurMannonov/medium_api_gateway/api/models"
	pbp "github.com/TemurMannonov/medium_api_gateway/genproto/post_service"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		req models.CreatePostRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.PostService().Create(context.Background(), &pbp.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      1,
		CategoryId:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post := parsePostModel(resp)
	c.JSON(http.StatusCreated, post)
}

func parsePostModel(post *pbp.Post) models.Post {
	return models.Post{
		ID:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserId,
		CategoryID:  post.CategoryId,
		CreatedAt:   post.CreatedAt,
		ViewsCount:  post.ViewsCount,
	}
}
