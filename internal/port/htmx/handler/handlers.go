package handler

import (
	"io/fs"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
)

type Handlers struct {
	Error Error
	Sign  Sign
	Home  Home
}

func NewHandlers(
	templateFS fs.FS,
	usecases application.UseCases,
) (Handlers, error) {
	errorHandler, err := NewErrorHandler(templateFS)
	if err != nil {
		return Handlers{}, err
	}
	signHandler, err := NewSign(templateFS, usecases.SignUseCase)
	if err != nil {
		return Handlers{}, err
	}
	homeHandler, err := NewHome(templateFS)
	if err != nil {
		return Handlers{}, err
	}

	return Handlers{
		Error: errorHandler,
		Sign:  signHandler,
		Home:  homeHandler,
	}, nil
}
