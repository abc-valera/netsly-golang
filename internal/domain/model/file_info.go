package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
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

type FileInfos []FileInfo

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

func FileTypeIsText(t FileType) bool {
	return t == FileTypeTXT
}

func FileTypeIsImage(t FileType) bool {
	return t == FileTypePNG || t == FileTypeJPEG
}

func FileTypeIsAudio(t FileType) bool {
	return t == FileTypeMP3
}

func FileTypeIsVideo(t FileType) bool {
	return t == FileTypeMP4
}
