package main

import (
	"fmt"
	"net/http"
	"os"
	"word-search-in-files/pkg/internal/response"
	"word-search-in-files/pkg/searcher"
)

func main() {
	http.HandleFunc("/files/search", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		searchValue := queryParams.Get("search")

		if searchValue == "" {
			response.SendError(w, http.StatusBadRequest, "notify search option")
			return
		}

		s := searcher.Searcher{FS: os.DirFS(".")}

		files, err := s.Search(searchValue)
		if err != nil {
			response.SendError(w, http.StatusBadRequest, "server error")
			return
		}

		response.SendData(w, http.StatusOK, files)
	})

	fmt.Printf("Server started on port: 8080")
	http.ListenAndServe(":8080", nil)
}
