package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Snippets struct {
	FilePath    string
	MatchLength int
	Matches     []int
}

type Search struct {
	Workdir  string
	log      Log
	articles []string
}

func (s *Search) init() {
	s.log.init()
	s.index()
}

func (s *Search) index() {
	err := filepath.Walk(s.Workdir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}
		s.articles = append(s.articles, path)
		return nil
	})
	if err != nil {
		s.log.Error(err, false)
	}
	s.log.Info("Successfully indexed directory")
}

func (s *Search) search(searchQuery []byte) {

}
