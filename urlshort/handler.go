package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if redirectURL, ok := pathsToUrls[req.URL.Path]; ok {
			http.Redirect(w, req, redirectURL, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, req)
	}
	// return fallback.ServeHTTP
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var PathUrls []PathUrl
	if err := yaml.Unmarshal(yml, &PathUrls); err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pathUrl := range PathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}

type PathUrlJSON struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var PathUrls []PathUrlJSON
	if err := json.Unmarshal(jsonData, &PathUrls); err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pathUrl := range PathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}
