package stato

import (
	"github.com/chrootlogin/stato/pkg/project"
	log "github.com/sirupsen/logrus"
)

type Stato struct {
	project project.Project
}

func (s *Stato) BuildAll() {
	log.Info("Generating all pages...")

	s.project.BuildAll()
}

func Load(workDir string) *Stato {
	p := project.Load(workDir)

	return &Stato{
		project: p,
	}
}