package model

import "github.com/abc-valera/netsly-golang/internal/domain/util/coderr"

var ErrFileContentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "File content not found")

type FileContent struct {
	ID      string
	Content []byte
}
