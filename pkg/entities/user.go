package entities

import (
	"github.com/ayhamal/gogql-boilerplate/graph/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName             string     `json:"full_name"`
	Email                string     `json:"email" gorm:"unique"`
	Role                 string     `json:"role"`
	Gender               string     `json:"gender"`
	Ocupation            string     `json:"ocupation"`
	ProfileImgUrl        string     `json:"profile_img_url"`
	IdentificationType   string     `json:"identification_type"`
	IdentificationNumber string     `json:"identification_number"`
	CountryCode          string     `json:"country_code"`
	PhoneNumber          string     `json:"phone_number"`
	Password             string     `json:"password"`
	Weight               *float64   `json:"weight"`
	Height               *float64   `json:"height"`
	Birthday             *time.Time `json:"birthday"`
}

// O(1)  To Graphql format
func (u *User) ToGql() *model.User {
	return &model.User{
		ID:        int(u.ID),
		FullName:  u.FullName,
		Email:     u.Email,
		Role:      model.Role(u.Role),
		Gender:    model.Gender(u.Gender),
		Ocupation: u.Ocupation,
		IdentificationType: model.IdentificationType(u.IdentificationType),
		IdentificationNumber: u.IdentificationNumber,
		CountryCode: u.CountryCode,
		PhoneNumber: u.PhoneNumber,
		Weight:    u.Weight,
		Height:    u.Height,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
