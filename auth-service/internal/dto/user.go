package dto

import "github.com/1abobik1/Cloud-Storage/internal/models"

type UserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *UserDTO) ToModel() *models.UserModel {
	return &models.UserModel{
		Email:    u.Email,
		Password: u.Password,
	}
}
