package service

import (
	"errors"
	"fmt"
	"gotpservice/internal/repository"
	"gotpservice/pkg/utils"
	"math/rand"
)

type OTPService interface {
    GenerateOTP(phone string) (string, error)
    VerifyOTP(phone, code string) (string, error)
}

type otpService struct {
    otpRepo  repository.OTPRepository
    userRepo repository.UserRepository
}

func NewOTPService(otpRepo repository.OTPRepository, userRepo repository.UserRepository) OTPService {
    return &otpService{
        otpRepo:  otpRepo,
        userRepo: userRepo,
    }
}

func (s *otpService) GenerateOTP(phone string) (string, error) {
    if !s.otpRepo.CanRequestOTP(phone) {
        return "", errors.New("rate limit exceeded")
    }

    otp := fmt.Sprintf("%06d", rand.Intn(1000000))
    fmt.Printf("OTP for %s: %s\n", phone, otp)

    return otp, s.otpRepo.SaveOTP(phone, otp)
}


func (s *otpService) VerifyOTP(phone, code string) (string, error) {
    if !s.otpRepo.ValidateOTP(phone, code) {
        return "", errors.New("invalid or expired OTP")
    }

    user, err := s.userRepo.GetByPhone(phone)
    if err != nil {
        user, err = s.userRepo.CreateUser(phone)
        if err != nil {
            return "", err
        }
    }

    token, err := utils.GenerateJWT(user.ID, phone)
    if err != nil {
        return "", err
    }

    return token, nil
}
