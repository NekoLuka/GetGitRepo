package main

import (
	"bytes"
	"crypto/rand"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Snippet struct {
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

func (s *Search) search(searchQuery []byte) []Snippet {
	var snippets []Snippet

	// Loop over the files indexed by the index function
	for _, v := range s.articles {
		file, err := os.Open(v)
		if err != nil {
			s.log.Error(err, false)
			file.Close()
			continue
		}
		info, err := file.Stat()
		if err != nil {
			s.log.Error(err, false)
			file.Close()
			continue
		}

		// Create the byte array to save the bytes
		data := make([]byte, info.Size())
		_, err = file.Read(data)
		if err != nil {
			s.log.Error(err, false)
			file.Close()
			continue
		}

		if !bytes.Contains(data, searchQuery) {
			file.Close()
			continue
		}

		// Loop through the file while saving the location of matches
		snippet := Snippet{FilePath: v, MatchLength: len(searchQuery)}
		index := 0
		replacement := make([]byte, len(searchQuery))
		_, _ = rand.Read(replacement)
		for true {
			index = bytes.Index(data, searchQuery)
			if index == -1 {
				break
			}
			snippet.Matches = append(snippet.Matches, index)
			data = bytes.Replace(data, searchQuery, replacement, 1)
		}

		snippets = append(snippets, snippet)
		file.Close()
	}
	return snippets
}
