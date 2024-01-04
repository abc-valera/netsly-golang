package common

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

type ITemplate interface {
	Render(w io.Writer, data interface{}) error
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
		return prodTemplate{}, codeerr.NewInternal("NewTemplate", fmt.Errorf("no filenames provided"))
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
		return prodTemplate{}, codeerr.NewInternal("NewTemplate", err)
	}

	// Get executeName
	executeName := filenames[0]
	if len(filenames) > 1 {
		executeName = strings.Split(filenames[len(filenames)-1], "/")[len(strings.Split(filenames[len(filenames)-1], "/"))-1]
	}

	return prodTemplate{
		executeName: executeName,
		tmpl:        t,
	}, nil
}

// Render executes the template with the given data.
func (t prodTemplate) Render(wr io.Writer, data interface{}) error {
	if t.tmpl == nil {
		return codeerr.NewInternal("Template.Render", fmt.Errorf("template is nil"))
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
		return prodTemplate{}, codeerr.NewInternal("NewTemplate", fmt.Errorf("no filenames provided"))
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

func (t devTemplate) Render(wr io.Writer, data interface{}) error {
	tmpl, err := template.ParseFS(t.fs, t.filenames...)
	if err != nil {
		return codeerr.NewInternal("Template.Render", err)
	}
	// Get executeName
	executeName := t.filenames[0]
	if len(t.filenames) > 1 {
		executeName = strings.Split(t.filenames[len(t.filenames)-1], "/")[len(strings.Split(t.filenames[len(t.filenames)-1], "/"))-1]
	}
	return tmpl.ExecuteTemplate(wr, executeName, data)
}

type Templates map[string]ITemplate

// NewTemplates creates a new map of Template instances.
// It accepts a slice of filenames.
func NewTemplates(isProd bool, fs fs.FS, filenamesSlices ...[]string) (Templates, error) {
	templates := make(map[string]ITemplate)
	for _, filenames := range filenamesSlices {
		if templates[filenames[0]] != nil {
			return nil, codeerr.NewInternal("NewTemplates", fmt.Errorf("template %s already exists", filenames[0]))
		}
		var (
			t   ITemplate
			err error
		)
		if isProd {
			t, err = newProdTemplate(fs, filenames...)
		} else {
			t, err = newDevTemplate(fs, filenames...)
		}
		if err != nil {
			return nil, err
		}
		templates[filenames[0]] = t
	}
	return templates, nil
}

// Render executes the template by its name with the given data.
func (t Templates) Render(wr io.Writer, name string, data interface{}) error {
	if !strings.HasSuffix(name, ".html") {
		name = name + ".html"
	}
	if t[name] == nil {
		return codeerr.NewInternal("Templates.Render", fmt.Errorf("template %s not found", name))
	}
	return t[name].Render(wr, data)
}
