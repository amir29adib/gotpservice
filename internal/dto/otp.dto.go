package dto

type OTPRequestDTO struct {
    Phone string `json:"phone" binding:"required,e164"`
}

type OTPVerifyDTO struct {
    Phone string `json:"phone" binding:"required,e164"`
    Code  string `json:"code" binding:"required,len=6"`
}
