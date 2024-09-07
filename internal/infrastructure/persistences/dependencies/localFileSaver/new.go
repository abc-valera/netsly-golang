package localFileSaver

import (
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

func New(localFileSaverFolderPath string) (string, error) {
	// Create the folder
	if err := os.MkdirAll(localFileSaverFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			return "", coderr.NewInternalErr(err)
		}
	}

	return localFileSaverFolderPath, nil
}
