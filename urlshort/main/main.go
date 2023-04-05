package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nekidb/gophercises/urlshort"
)

func main() {
	fileName := flag.String("file", "paths_to_urls.yaml", "name of file with yaml data")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	//
	// 	// Build the YAMLHandler using the mapHandler as the
	// 	// fallback

	jsonData := `[{"path":"/json-godoc","url":"https://pkg.go.dev/encoding/json"}]`
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonData), mapHandler)
	if err != nil {
		panic(err)
	}

	yamlData, err := getFileData(*fileName)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), jsonHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func getFileData(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fileData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
