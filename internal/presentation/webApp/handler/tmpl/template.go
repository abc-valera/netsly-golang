package tmpl

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
)

type Data map[string]interface{}

type ITemplate interface {
	Render(w http.ResponseWriter, data interface{}) error
}

// prodTemplate is a wrapper for template.prodTemplate.
type prodTemplate struct {
	executeName string             // executeName is the name of the template to be executed through Execute() method
	tmpl        *template.Template // t is the template.Template instance
}

// newProdTemplate creates a new Template instance.
// Allows to pass filenames without '.html' extension.
// Filenames are passed in the order from the most specific to the most general.
func newProdTemplate(fs fs.FS, filenames ...string) (ITemplate, error) {
	// Check if filenames is empty
	if len(filenames) == 0 {
		return prodTemplate{}, coderr.NewInternalErr(fmt.Errorf("no filenames provided"))
	}
	// Add .html extension if not present (this allows to pass filenames without extension)
	for i, filename := range filenames {
		if !strings.HasSuffix(filename, ".html") {
			filenames[i] = filename + ".html"
		}
	}

	// Parse template
	t, err := template.ParseFS(fs, filenames...)
	if err != nil {
		return prodTemplate{}, coderr.NewInternalErr(err)
	}

	// Get executeName
	executeName := strings.Split(filenames[len(filenames)-1], "/")[len(strings.Split(filenames[len(filenames)-1], "/"))-1]

	return prodTemplate{
		executeName: executeName,
		tmpl:        t,
	}, nil
}

// Render executes the template with the given data.
func (t prodTemplate) Render(wr http.ResponseWriter, data interface{}) error {
	if t.tmpl == nil {
		return coderr.NewInternalErr(fmt.Errorf("template is nil"))
	}

	return t.tmpl.ExecuteTemplate(wr, t.executeName, data)
}

// devTemplate is a wrapper for template.prodTemplate.
// It allows to change the template without restarting the server.
type devTemplate struct {
	fs        fs.FS
	filenames []string
}

// newDevTemplate creates a new Template instance.
// Allows to pass filenames without '.html' extension.
// Filenames are passed in the order from the most specific to the most general.
func newDevTemplate(fs fs.FS, filenames ...string) (ITemplate, error) {
	// Check if filenames is empty
	if len(filenames) == 0 {
		return prodTemplate{}, coderr.NewInternalErr(fmt.Errorf("no filenames provided"))
	}
	// Add .html extension if not present (this allows to pass filenames without extension)
	for i, filename := range filenames {
		if !strings.HasSuffix(filename, ".html") {
			filenames[i] = filename + ".html"
		}
	}
	return devTemplate{
		fs:        fs,
		filenames: filenames,
	}, nil
}

func (t devTemplate) Render(wr http.ResponseWriter, data interface{}) error {
	tmpl, err := template.ParseFS(t.fs, t.filenames...)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	// Get executeName
	executeName := strings.Split(t.filenames[len(t.filenames)-1], "/")[len(strings.Split(t.filenames[len(t.filenames)-1], "/"))-1]

	return tmpl.ExecuteTemplate(wr, executeName, data)
}

func NewTemplate(fs fs.FS, filenames ...string) (ITemplate, error) {
	if global.Mode() == mode.Development {
		return newDevTemplate(fs, filenames...)
	}
	return newProdTemplate(fs, filenames...)
}
