package localgaesupport

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type AppYamlHandler struct {
	URL         string            `yaml:"url"`
	StaticFiles string            `yaml:"static_files"`
	StaticDir   string            `yaml:"static_dir"`
	MimeType    string            `yaml:"mime_type"`
	HTTPHeaders map[string]string `yaml:"http_headers"`

	pattern *regexp.Regexp `yaml:"-"`
}

func (h *AppYamlHandler) Setup() error {
	s := `\A` + h.StaticFiles + `\z`
	re, err := regexp.Compile(s)
	if err != nil {
		log.Printf("Invalid regexp %v because of %v", s, err)
		return err
	}
	h.pattern = re
	return nil
}

var slashNumPattern = regexp.MustCompile(`\\(\d+)`)

func (h *AppYamlHandler) BuildPath(w http.ResponseWriter, r *http.Request, m1 []string) (string, error) {
	var rerr error
	path := slashNumPattern.ReplaceAllStringFunc(h.StaticFiles, func(m2 string) string {
		index, err := strconv.Atoi(m2[1:])
		if err != nil {
			rerr = fmt.Errorf("Failed to parse slash number %s because of %v\n", m2[1:], err)
			return ""
		}
		if index <= 0 || index >= len(m1) {
			rerr = fmt.Errorf("Invalid index %s must be >= 0 and must be < %d\n", len(m1))
			return ""
		}
		return m1[index]
	})
	if rerr != nil {
		return "", rerr
	}
	return path, nil
}

func (h *AppYamlHandler) ProcessHeaders(w http.ResponseWriter, r *http.Request) {
	if h.MimeType != "" {
		w.Header().Set("Content-Type", h.MimeType)
	}
	for k, v := range h.HTTPHeaders {
		w.Header().Set(k, v)
	}
}
