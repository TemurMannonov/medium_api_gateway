package v1

import (
	"context"
	"net/http"

	"github.com/TemurMannonov/medium_api_gateway/api/models"
	pbu "github.com/TemurMannonov/medium_api_gateway/genproto/user_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Router /auth/register [post]
// @Summary Register a user
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Register(c *gin.Context) {
	var (
		req models.RegisterRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, _ := h.grpcClient.UserService().GetByEmail(context.Background(), &pbu.GetByEmailRequest{
		Email: req.Email,
	})
	if user != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrEmailExists))
		return
	}

	_, err = h.grpcClient.AuthService().Register(context.Background(), &pbu.RegisterRequest{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "success",
	})
}

// @Router /auth/verify [post]
// @Summary Verify user
// @Description Verify user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Verify(c *gin.Context) {
	var (
		req models.VerifyRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.AuthService().Verify(context.Background(), &pbu.VerifyRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == "incorrect_code" {
			c.JSON(http.StatusBadRequest, errorResponse(ErrIncorrectCode))
			return
		} else if s.Message() == "code_expired" {
			c.JSON(http.StatusBadRequest, errorResponse(ErrCodeExpired))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	c.JSON(http.StatusCreated, models.AuthResponse{
		ID:          result.Id,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: result.AccessToken,
	})
}

// @Router /auth/login [post]
// @Summary Login User
// @Description Login User
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Login(c *gin.Context) {
	var (
		req models.LoginRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.logger.WithError(err).Error("failed to bind JSON in login")
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.grpcClient.AuthService().Login(context.Background(), &pbu.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to login")
		s, _ := status.FromError(err)
		if s.Code() == codes.NotFound || s.Message() == "incorrect_password" {
			c.JSON(http.StatusBadRequest, errorResponse(ErrWrongEmailOrPass))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		ID:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: user.AccessToken,
	})

}
