package theme

import (
	"bytes"
	"errors"
	"github.com/chrootlogin/stato/pkg/models"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"text/template"
	"time"
)

// public interface
type ThemeEngine interface {
	Render(layoutName string, data *models.ViewData) (*bytes.Buffer, error)
}

// private struct
type themeEngine struct {
	cache *cache.Cache
	themeDir string
}

func New(themeDir string) ThemeEngine {
	return &themeEngine {
		cache: cache.New(5 * time.Minute, 10 * time.Minute),
		themeDir: themeDir,
	}
}

func (t *themeEngine) Render(layoutName string, data *models.ViewData) (*bytes.Buffer, error) {
	var tmpl *template.Template

	tpl, found := t.cache.Get("layout:" + layoutName)
	if found {
		tmpl = tpl.(*template.Template)
	} else {
		// build layout
		tmpl = t.buildLayout(layoutName)

		// write to cache
		t.cache.SetDefault("layout:" + layoutName, tmpl)
	}

	// create buffer
	buf := new(bytes.Buffer)

	// execute template
	err := tmpl.ExecuteTemplate(buf, layoutName, data)
	if err != nil {
		return nil, errors.New("error executing template")
	}

	return buf, nil
}

func (t *themeEngine) buildLayout(layoutName string) *template.Template {
	// load layout
	tmpl, err := template.ParseFiles(filepath.Join(t.themeDir, "layouts", layoutName));
	if err != nil {
		log.Fatal("Error loading layout", err)
	}

	// load templates
	tplGlob := filepath.Join(t.themeDir, "templates", "*.tpl")
	_, err = tmpl.ParseGlob(tplGlob);
	if err != nil {
		log.Fatal("Error loading layout", err)
	}

	return tmpl
}