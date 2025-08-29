package model

import "time"

type User struct {
    ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Phone           string    `gorm:"uniqueIndex;size:20" json:"phone"`
    RegistrationDate time.Time `json:"registration_date"`
}
