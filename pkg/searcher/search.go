package searcher

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"sync"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

func (s *Searcher) Search(word string) ([]string, error) {
	osFiles, err := dir.FilesFS(s.FS, ".")
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var containWordFiles []string
	for _, file := range osFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if containsWord(fmt.Sprintf("./examples/%s", file), word) {
				containWordFiles = append(containWordFiles, file)
			}
		}(file)
	}

	wg.Wait()

	return containWordFiles, nil
}

func containsWord(filePath string, word string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), word) {
			return true
		}
	}
	return false
}
