package model

import "github.com/abc-valera/netsly-golang/internal/core/coderr"

var ErrFileContentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "File content not found")

type FileContent []byte

type FileContents []FileContent
