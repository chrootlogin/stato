package project

import (
	"bufio"
	"bytes"
	"github.com/chrootlogin/stato/pkg/models"
	"github.com/chrootlogin/stato/pkg/renderer"
	"github.com/chrootlogin/stato/pkg/theme"
	"github.com/chrootlogin/stato/pkg/utils/consts"
	"github.com/chrootlogin/stato/pkg/utils/helper"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Project interface {
	BuildAll()
	RenderFile(path string) (*bytes.Buffer, error)
}

type project struct {
	config *config
	workDir string
	cache *cache.Cache
	theme theme.ThemeEngine
}

func Load(workDir string) Project {
	p := &project{
		config: &config{},
		cache: cache.New(5 * time.Minute, 10 * time.Minute),
		workDir: workDir,
	}

	// init project
	p.init()

	return p
}

func (p *project) RenderFile(path string) (*bytes.Buffer, error) {
	log.WithField("path", path).Info("rendering file")

	// create view data object
	data := &models.ViewData{
		Site: models.ViewSiteData{
			Title: p.config.Title,
			LanguageCode: p.config.LanguageCode,
		},
		Layout: consts.StatoDefaultLayoutFile,
	}

	// render template file
	err := renderer.Render(path, data)
	if err != nil {
		return nil, err
	}

	// render layout with template
	buf, err := p.theme.Render("single.tpl", data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (p *project) BuildTemplate(path string) {
	buf, err := p.RenderFile(path)
	if err != nil {
		log.WithField("path", path).Warn(err)
		return
	}

	// get output path
	outputPath := p.buildOutputPath(path)

	// create output dir
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.WithField("path", path).Error("error creating output path", err)
		return
	}

	// create file
	file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.WithField("path", path).Error("error creating output file", err)
		return
	}

	// write file
	writer := bufio.NewWriter(file)
	_, err = writer.Write(buf.Bytes())
	if err != nil {
		log.WithField("path", path).Error("error writing to output file", err)
		return
	}

	// flush buffer to file
	if err := writer.Flush(); err != nil {
		log.WithField("path", path).Error("error flushing buffer", err)
	}

	// close file
	if err := file.Close(); err != nil {
		log.WithField("path", path).Error("error closing output file", err)
	}
}

func (p *project) BuildAll() {
	err := filepath.Walk(filepath.Join(p.workDir, consts.StatoContentPath),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// ignore directories
			if info.IsDir() {
				return nil;
			}

			p.BuildTemplate(path)

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	p.CopyStaticFiles()
}

func (p *project) buildOutputPath(path string) string {
	// build new path
	outputPath := strings.TrimPrefix(path, filepath.Join(p.workDir, consts.StatoContentPath))
	fileName := filepath.Base(outputPath)

	// get filename
	outputPath = strings.TrimSuffix(outputPath, fileName)
	outputName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	// put every page in subdirectory to have pretty urls
	if outputName == "index" {
		// build new filepath
		outputPath = filepath.Join(p.workDir, consts.StatoOutputPath, outputPath, "index.html")
	} else {
		outputPath = filepath.Join(p.workDir, consts.StatoOutputPath, outputPath, outputName, "index.html")
	}

	return outputPath
}

func (p *project) CopyStaticFiles() {
	staticPath := filepath.Join(p.workDir, consts.StatoStaticPath)
	staticThemePath := filepath.Join(p.workDir, consts.StatoThemesPath, p.config.Theme, consts.StatoStaticPath)
	outputPath := filepath.Join(p.workDir, consts.StatoOutputPath)

	if err := helper.CopyDir(staticPath, outputPath); err != nil {
		log.WithFields(log.Fields{
			"src": staticPath,
			"dst": outputPath,
		}).Error("error copying static files", err)
	}

	if err := helper.CopyDir(staticThemePath, outputPath); err != nil {
		log.WithFields(log.Fields{
			"src": staticPath,
			"dst": outputPath,
		}).Error("error copying static files", err)
	}
}

func (p *project) init() {
	p.initConfig()
	p.initTheme()
}

func (p *project) initConfig() {
	configFilePath := filepath.Join(p.workDir, consts.StatoDefaultCfgFile)
	log.WithField("path", configFilePath).Info("loading project config")

	// read file
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.WithField("path", configFilePath).Fatal("error loading project config")
	}

	// unmarshal config
	err = yaml.Unmarshal(yamlFile, p.config)
	if err != nil {
		log.Fatal("error unmarshalling project config", err)
	}
}

func (p *project) initTheme() {
	themeDir := filepath.Join(p.workDir, "themes", p.config.Theme)

	p.theme = theme.New(themeDir)
}