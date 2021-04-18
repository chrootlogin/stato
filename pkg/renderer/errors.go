package renderer

import "errors"

var (
	ErrNoFileSeparator       = errors.New("error no file separator")
	ErrNoSuitableRenderer    = errors.New("found no suitable renderer for template")
	ErrParsingTemplateFile   = errors.New("error parsing template file")
	ErrRenderingTemplateFile = errors.New("error rendering template file")
)
