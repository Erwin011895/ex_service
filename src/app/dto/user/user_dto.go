package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserReqDTOInterface interface {
	Validate() error
}

type RegisterReqDTO struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *RegisterReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
		validation.Field(&dto.Email, validation.Required),
		validation.Field(&dto.Password, validation.Required),
		validation.Field(&dto.Password, validation.Length(8, 30)),
	); err != nil {
		return err
	}
	return nil
}

type RegisterRespDTO struct {
	ID       int64  `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	WalletID int64  `json:"wallet_id" db:"wallet_id"`
	Token    string `json:"token"`
}

type LoginReqDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}
