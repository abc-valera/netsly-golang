package restHandler

import (
	"fmt"
	"net/http"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
)

type Docs struct {
	openapiFile []byte
}

func NewDocs(openApiFile []byte) Docs {
	return Docs{
		openapiFile: openApiFile,
	}
}

func (d Docs) OpenApiFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/yaml")
	w.Write(d.openapiFile)
}

func (Docs) ScalarDocs(w http.ResponseWriter, r *http.Request) {
	scalarOpenapiVisualizer := `
		<!doctype html>
		<html>
		<head>
			<title>Scalar API Reference</title>
			<meta charset="utf-8" />
			<meta
			name="viewport"
			content="width=device-width, initial-scale=1" />
		</head>
		<body>
			<script
			id="api-reference"
			data-url="%s"></script>
			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		</body>
		</html>
	`

	https := "http"
	if global.IsHTTPS() {
		https = "https"
	}

	// TODO: come up with a way of passing link here
	openapiUrl := https + "://" + global.SubdomainJsonApi() + global.DomainName() + "/v1/openapi"

	scalarOpenapiVisualizer = fmt.Sprintf(scalarOpenapiVisualizer, openapiUrl)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(scalarOpenapiVisualizer))
}
