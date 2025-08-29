package dto

type UserDTO struct {
    ID uint `json:"id"`
    Phone string `json:"phone"`
    RegistrationDate string `json:"registration_date"`
}

type UserListResponseDTO struct {
    Users []UserDTO `json:"users"`
    Total int       `json:"total"`
}