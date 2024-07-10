package searcher

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

func (s *Searcher) Search(word string) (files []string, err error) {
	osFiles, err := dir.FilesFS(s.FS, ".")
	var containWordFiles []string

	if err != nil {
		return nil, err
	}

	for _, file := range osFiles {
		filePath := fmt.Sprintf("./examples/%s", file)
		contain := isContainsWord(filePath, word)
		if contain {
			containWordFiles = append(containWordFiles, file)
		}
	}

	return containWordFiles, nil
}

func isContainsWord(filePath string, word string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileWords := strings.Fields(line)

		for _, fileWord := range fileWords {
			if word == fileWord {
				return true
			}
		}
	}
	return false
}
