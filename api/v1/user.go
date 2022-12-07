package v1

import (
	"context"
	"net/http"

	"github.com/TemurMannonov/medium_api_gateway/api/models"
	"github.com/gin-gonic/gin"

	pbu "github.com/TemurMannonov/medium_api_gateway/genproto/user_service"
)

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req models.CreateUserRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.grpcClient.UserService().Create(context.Background(), &pbu.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     *req.PhoneNumber,
		Email:           req.Email,
		Gender:          *req.Gender,
		Password:        req.Password,
		Username:        *req.Username,
		ProfileImageUrl: *req.ProfileImageUrl,
		Type:            req.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.User{
		ID:              user.Id,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PhoneNumber:     &user.PhoneNumber,
		Email:           user.Email,
		Gender:          &user.Gender,
		Username:        &user.Username,
		ProfileImageUrl: &user.ProfileImageUrl,
		Type:            user.Type,
		CreatedAt:       user.CreatedAt,
	})
}
