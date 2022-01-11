package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Failed string
	Tag    string
	Value  interface{}
}

func ErrorHandler(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validated := validator.New()
	err := validated.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var ep ErrorResponse
			ep.Failed = e.StructNamespace()
			ep.Tag = e.Tag()
			ep.Value = e.Param()
			errors = append(errors, &ep)
		}
	}
	return errors
}

func BuildResponse(m interface{}, s bool) interface{} {
	return fiber.Map{
		"message": m,
		"status": s,
	}
}