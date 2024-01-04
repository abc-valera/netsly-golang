package handler

import (
	"io/fs"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
)

type Handlers struct {
	ErrorHandler ErrorHandler
	SignHandler  SignHandler
}

func NewHandlers(
	templateFS fs.FS,
	usecases application.UseCases,
) (Handlers, error) {
	errorHandler, err := NewErrorHandler(templateFS)
	if err != nil {
		return Handlers{}, err
	}
	signHandler, err := NewSignHandler(templateFS, usecases.SignUseCase)
	if err != nil {
		return Handlers{}, err
	}

	return Handlers{
		ErrorHandler: errorHandler,
		SignHandler:  signHandler,
	}, nil
}
