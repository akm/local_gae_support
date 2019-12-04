package localgaesupport

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AppYamlHandlers []*AppYamlHandler

func (s AppYamlHandlers) Each(f func(*AppYamlHandler) error) error {
	for _, i := range s {
		if err := f(i); err != nil {
			return err
		}
	}
	return nil
}

func (s AppYamlHandlers) Select(f func(*AppYamlHandler) bool) AppYamlHandlers {
	r := AppYamlHandlers{}
	for _, i := range s {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (s AppYamlHandlers) NewHandler(defaultHandler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("AppYamlHandlers: %v %q\n", r.Method, r.URL.Path)
		if r.Method == http.MethodGet {
			for idx, i := range s {
				log.Printf("AppYamlHandlers[%d]: %s\n", idx, i.URL)
				m := i.pattern.FindStringSubmatch(r.URL.Path)
				if m == nil {
					continue
				}
				path, err := i.BuildPath(w, r, m)
				if err != nil {
					log.Printf("WARNING BuildPath returned an error: %v\n", err)
					continue
				}
				log.Printf("AppYamlHandlers[%d]: %s ==> %s\n", idx, i.URL, path)
				i.ProcessHeaders(w, r)

				if b, err := ioutil.ReadFile(path); err != nil {
					log.Printf("Failed to read file to return because of %v", err)
					continue
				} else {
					// http.ServeFile(w, r, path)
					log.Printf("Returning static file %s size %d bytes", path, len(b))
					fi, err := os.Stat(path)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					} else {
						http.ServeContent(w, r, path, fi.ModTime(), bytes.NewReader(b))
					}
				}

				return
			}
		}
		defaultHandler.ServeHTTP(w, r)
	}
}
