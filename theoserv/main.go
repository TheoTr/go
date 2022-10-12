package main

import (
	"net/http"

	map_struct "github.com/TheoTr/go/map"
)

func main() {
	dict := map_struct.Dictionary{"test": "this is just a test"}
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		queryStr := r.URL.Query().Get("query")
		searchResult, err := dict.Search(queryStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(searchResult))
	})

	http.ListenAndServe("localhost:7777", nil)
}

//func ResponseWritter(){}
// localhost:7777/search?query=test