package renderer

import (
	"github.com/chrootlogin/stato/pkg/models"
	"github.com/chrootlogin/stato/pkg/utils/consts"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type FileRenderer interface {
	Render(data *models.ViewData) error
	IsSupported(path string) bool
}

var frs []FileRenderer

func Render(path string, data *models.ViewData) error {
	for _, fileRenderer := range frs {
		if fileRenderer.IsSupported(path) {
			err := readFile(path, data)
			if err != nil {
				return err
			}

			return fileRenderer.Render(data)
		}
	}

	return ErrNoSuitableRenderer
}

func appendRenderer(fr FileRenderer) {
	frs = append(frs, fr)
}

func readFile(path string, data *models.ViewData) error {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// working with string
	data.Content = string(d)

	// check for config prefix
	if strings.HasPrefix(data.Content, consts.StatoTemplateHeader) {
		// remove the prefix
		data.Content = strings.TrimPrefix(data.Content, consts.StatoTemplateHeader)

		// split data to extract config and content
		dataSplt := strings.SplitN(data.Content, consts.StatoTemplateSeparator, 2)
		if len(dataSplt) != 2 {
			return ErrNoFileSeparator
		}

		data.Content = dataSplt[1]  // content
		dataCfg := dataSplt[0] // config

		// unmarshal config
		var templateConfig models.TemplateConfig
		err := yaml.Unmarshal([]byte(dataCfg), &templateConfig)
		if err != nil {
			return err
		}

		// set data from template
		data.Title = templateConfig.Title
		if templateConfig.Layout != "" {
			data.Layout = templateConfig.Layout
		}
	}

	return nil
}