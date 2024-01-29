package handler

import (
	"io/fs"

	"github.com/abc-valera/netsly-api-golang/internal/application"
)

type Handlers struct {
	Error Error
	Sign  Sign
	Home  Home
}

func NewHandlers(
	templateFS fs.FS,
	usecases application.UseCases,
) Handlers {
	return Handlers{
		Error: NewErrorHandler(templateFS),
		Sign:  NewSign(templateFS, usecases.SignUseCase),
		Home:  NewHome(templateFS),
	}
}
