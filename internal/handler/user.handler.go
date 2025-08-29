package handler

import (
	_ "gotpservice/internal/dto"
	"gotpservice/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{userService: s}
}

// GetUserByPhone godoc
// @Summary Get user by phone number
// @Description Retrieve a user by their phone number
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param phone path string true "Phone Number"
// @Success 200 {object} dto.UserDTO
// @Failure 404 {object} map[string]string
// @Router /users/{phone} [get]
func (h *UserHandler) GetUserByPhone(c *gin.Context) {
    phoneParam := c.Param("phone")
    phoneFromToken, exists := c.Get("phone")
    if !exists || phoneParam != phoneFromToken {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only access your own profile"})
        return
    }

    user, err := h.userService.GetByPhone(phoneParam)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of users with pagination and optional search by phone
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param search query string false "Search by phone"
// @Success 200 {object} dto.UserListResponseDTO
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    search := c.Query("search")

    users, total := h.userService.ListUsers(page, limit, search)
    c.JSON(http.StatusOK, gin.H{
        "total": total,
        "users": users,
    })
}
