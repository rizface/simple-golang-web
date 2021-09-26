package helper

import (
	"errors"
	"pbl-orkom/app"

	"github.com/go-playground/validator/v10"
)

func ValidationHandler(err error) {
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := err.Field() + app.Message[err.Tag()]
			panic(errors.New(msg))
		}
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
