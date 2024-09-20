package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func TestFileCreate(t *testing.T) {
	ctx, r, entities := NewTest(t)

	resp, err := entities.File.Create(ctx, entity.FileCreateRequest{
		Name:        "testName",
		Type:        model.FileTypeTXT,
		FileContent: []byte("Test File Content"),
	})
	r.NoError(err)
	r.NotEmpty(resp)

	actualFileInfo, actualFileContent, err := entities.File.GetByID(ctx, resp.FileInfo.ID)
	r.NoError(err)
	r.NotEmpty(actualFileInfo)
	r.NotEmpty(actualFileContent)

	r.Equal(resp.FileInfo, actualFileInfo)
	r.Equal(resp.FileContent, actualFileContent)
}

func TestFileUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestFileDelete(t *testing.T) {
	t.Skip("unimplemented")
}
