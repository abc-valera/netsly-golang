package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrFileInfoNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "File info not found")

type FileInfo struct {
	ID        string
	Name      string
	Type      FileType
	Size      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type FileType int

const (
	FileTypeTXT FileType = iota + 1

	FileTypePNG
	FileTypeJPEG

	FileTypeMP3

	FileTypeMP4
	FileTypeGIF

	fileTypeEnd
)

func (t FileType) IsValid() bool {
	return t > 0 && t < fileTypeEnd
}
