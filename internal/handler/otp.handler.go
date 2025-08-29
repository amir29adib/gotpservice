package handler

import (
	"gotpservice/internal/dto"
	"gotpservice/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
    otpService service.OTPService
}

func NewOTPHandler(s service.OTPService) *OTPHandler {
    return &OTPHandler{otpService: s}
}

// RequestOTP godoc
// @Summary Request OTP code
// @Description Generate a one-time password (OTP) and print to console
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.OTPRequestDTO true "Phone number payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 429 {object} map[string]string
// @Router /auth/request-otp [post]
func (h *OTPHandler) RequestOTP(c *gin.Context) {
    var req dto.OTPRequestDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    otp, err := h.otpService.GenerateOTP(req.Phone)
    if err != nil {
        c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "OTP generated (for dev)",
        "otp":     otp,
    })
}

// VerifyOTP godoc
// @Summary Verify OTP and login/register
// @Description Validate OTP, register new user if not exists, and return JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.OTPVerifyDTO true "Phone and OTP code"
// @Success 200 {object} dto.TokenResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/verify-otp [post]
func (h *OTPHandler) VerifyOTP(c *gin.Context) {
    var req dto.OTPVerifyDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.otpService.VerifyOTP(req.Phone, req.Code)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, dto.TokenResponseDTO{Token: token})
}
