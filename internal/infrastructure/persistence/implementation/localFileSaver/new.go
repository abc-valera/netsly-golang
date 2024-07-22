package localFileSaver

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

func New(filesPath string) (string, error) {
	if filesPath == "" {
		return "", coderr.NewInternalString("filesPath is empty")
	}

	return filesPath, nil
}
