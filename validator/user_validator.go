package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/shundahoy/techready/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスは必須項目です。"),
			validation.RuneLength(1, 30).Error("30文字以内で入力してください。"),
			is.Email.Error("フォーマットが正しくありません。"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードは必須項目です。"),
			validation.RuneLength(6, 30).Error("6~30文字以内で入力してください。"),
		),
	)
}
