package service

import (
	"gotpservice/internal/dto"
	"gotpservice/internal/repository"
)

type UserService interface {
    GetByPhone(phone string) (*dto.UserDTO, error)
    ListUsers(page, limit int, search string) ([]dto.UserDTO, int)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

func (s *userService) GetByPhone(phone string) (*dto.UserDTO, error) {
    user, err := s.userRepo.GetByPhone(phone)
    if err != nil {
        return nil, err
    }
    return &dto.UserDTO{
        ID: user.ID,
        Phone: user.Phone,
        RegistrationDate: user.RegistrationDate.Format("2006-01-02 15:04:05"),
    }, nil
}

func (s *userService) ListUsers(page int, limit int, search string) ([]dto.UserDTO, int) {
    users, total := s.userRepo.ListUsers(page, limit, search)

    var result []dto.UserDTO
    for _, u := range users {
        result = append(result, dto.UserDTO{
            ID: u.ID,
            Phone: u.Phone,
            RegistrationDate: u.RegistrationDate.Format("2006-01-02 15:04:05"),
        })
    }
    return result, total
}
