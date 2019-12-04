package localgaesupport

import (
	"log"
	"net/http"
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
		if r.Method == http.MethodGet {
			for _, i := range s {
				m := i.pattern.FindStringSubmatch(r.URL.Path)
				if m == nil {
					continue
				}
				path, err := i.BuildPath(w, r, m)
				if err != nil {
					log.Printf("WARNING BuildPath returned an error: %v\n", err)
					continue
				}
				i.ProcessHeaders(w, r)
				http.ServeFile(w, r, path)
				return
			}
		}
		defaultHandler.ServeHTTP(w, r)
	}
}
