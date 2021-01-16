package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// https://golang.org/pkg/net/http/#HandleFunc - a http.HandlerFunc is going to be returned here
	// which looks like this, and we don't need to typecast it like 
	// "return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){})"
	// because we specified the return type of MapHandler like so, so it will know that!
	return func(w http.ResponseWriter, r *http.Request) {
		// we only need the path of the urls from the request
		path := r.URL.Path
		// instead of 'err', 'ok' will be true if path found in our map, and what 
		// is found will be passed as the value of 'dest'
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			// we don't want our program to run further from here
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
