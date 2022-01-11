package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	GenderID int32     `json:"gender_id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type Gender struct {
	ID     int64  `json:"id"`
	Gender string `json:"gender"`
}

type ErrorResponse struct {
	Failed string
	Tag    string
	Value  string
}

func ValidateFilm(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	validated := validator.New()
	err := validated.Struct(user)
	if err != nil {
		for _, error := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Failed = error.StructNamespace()
			element.Tag = error.Tag()
			element.Value = error.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
