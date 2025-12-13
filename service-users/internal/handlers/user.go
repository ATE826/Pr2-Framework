package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SpiritFoxo/control-system-microservices/service-users/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func response(c *gin.Context, status int, success bool, data interface{}, err error) {
	if err != nil {
		c.JSON(status, gin.H{
			"success": success,
			"error": gin.H{
				"code":    status,
				"message": err.Error(),
			},
		})
		return
	}
	c.JSON(status, gin.H{
		"success": success,
		"data":    data,
	})
}

// RegisterUser
// RegisterUser
// @Summary Creates a new user
// @Description Creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body services.RegisterUserInput true "User data"
// @Success 201 {object} services.UserResponse "Successfully created user"
// @Security BearerAuth
// @Router /admin/users/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input services.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response(c, http.StatusBadRequest, false, nil, err)
		return
	}

	user, err := h.service.RegisterUser(input)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "email already exists" || err.Error() == "invalid email format" ||
			err.Error() == "password must be at least 8 characters" ||
			err.Error() == "email, password, and name are required" ||
			err.Error()[:12] == "invalid role" {
			status = http.StatusBadRequest
		}
		response(c, status, false, nil, err)
		return
	}

	response(c, http.StatusCreated, true, user, nil)
}

// LoginUser
// @Summary performs login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body services.LoginInput true "User data"
// @Succes 200 {object} map[string]interface{} "User auth data"
// @Router /auth/login [post]
func (h *UserHandler) LoginUser(c *gin.Context) {

	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response(c, http.StatusBadRequest, false, nil, err)
		return
	}

	token, user, err := h.service.LoginUser(input.Email, input.Password)
	if err != nil {
		response(c, http.StatusUnauthorized, false, nil, err)
		return
	}

	response(c, http.StatusOK, true, gin.H{
		"token": token,
		"user":  user,
	}, nil)
}

// GetUserById
// @Summary Gets user by ID
// @Description Gets user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} services.UserResponse "User data"
// @Security BearerAuth
// @Router /admin/users/{userId} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response(c, http.StatusBadRequest, false, nil, errors.New("invalid user ID"))
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		response(c, http.StatusNotFound, false, nil, err)
		return
	}

	response(c, http.StatusOK, true, user, nil)
}
