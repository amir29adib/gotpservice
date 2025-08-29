package repository

import (
	"gotpservice/internal/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
    GetByPhone(phone string) (*model.User, error)
    CreateUser(phone string) (*model.User, error)
	ListUsers(page int, limit int, search string) ([]model.User, int)
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) GetByPhone(phone string) (*model.User, error) {
    var user model.User
    if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) CreateUser(phone string) (*model.User, error) {
    user := model.User{
        Phone:           phone,
        RegistrationDate: time.Now(),
    }
    if err := r.db.Create(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) ListUsers(page, limit int, search string) ([]model.User, int) {
    var users []model.User
    var total int64

    query := r.db.Model(&model.User{})
    if search != "" {
        query = query.Where("phone ILIKE ?", "%"+search+"%")
    }

    query.Count(&total)

    query = query.Offset((page - 1) * limit).Limit(limit).Order("registration_date desc")
    query.Find(&users)

    return users, int(total)
}
