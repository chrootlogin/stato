package webserver

import (
	"github.com/chrootlogin/stato/pkg/project"
	"github.com/chrootlogin/stato/pkg/utils/helper"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Webserver struct {
	router *gin.Engine
	project project.Project
}

func New(workDir string) *Webserver {
	// set gin mode
	if log.GetLevel() == log.DebugLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// new gin router
	r := gin.New()

	// add logrus
	r.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true))

	return &Webserver{
		router: r,
		project: project.Load(workDir),
	}
}

func (w *Webserver) Router(c *gin.Context) {
	// extract path from url
	url := c.Request.URL.String()[1:]

	// try to build template
	buf, err := w.project.BuildFile(url)
	if err == nil {
		// read data from buffer
		c.DataFromReader(
			http.StatusOK,
			int64(buf.Len()),
			"text/html",
			buf,
			map[string]string{},
		)
		return
	}

	// try to get static file
	s, rd, err := w.project.GetStatic(url)
	if err == nil {
		// read data from buffer
		c.DataFromReader(
			http.StatusOK,
			s,
			helper.GetMimetype(url),
			rd,
			map[string]string{},
		)
		return
	}

	c.String(404, "not found")
}

func (w *Webserver) Run(listenAddress string) {
	w.router.NoRoute(w.Router)

	if err := w.router.Run(listenAddress); err != nil {
		log.Fatal(err)
	}
}