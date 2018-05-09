package main

import (
	"net/http"
	"io"
	"os"
	"encoding/json"
	"compress/gzip"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}
	gzipWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
