package organizer

import (
	"path/filepath"
	"regexp"
	"strings"
)

var videoTypes = []string{".mkv", ".mp4"}

type vidFile struct {
	Path string
}

func newVidFile(srcFile string) vidFile {
	p, err := filepath.Abs(srcFile)
	if err != nil {
		panic(err)
	}

	return vidFile{
		Path: p,
	}
}

func (v vidFile) translationGroup() string {
	r, err := regexp.Compile(`\[([^\]]*)\]`)
	if err != nil {
		return ""
	}

	return r.FindString(v.Path)
}

func (v vidFile) Title() string {
	s := strings.Replace(v.fileName(), v.translationGroup(), "", 1)
	s = strings.Replace(s, v.endSection(), "", 1)
	s = strings.TrimSpace(s)

	return s
}

func (v vidFile) endSection() string {
	r, err := regexp.Compile(`.*(-.*)`)
	if err != nil {
		return ""
	}

	matches := r.FindStringSubmatch(v.Path)

	lastIndex := len(matches) - 1
	return matches[lastIndex]
}

func (v vidFile) fileName() string {
	return filepath.Base(v.Path)
}
