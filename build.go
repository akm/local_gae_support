package localgaesupport

import (
	"net/http"
)

func Static(appYamlPath string, defaultHandler http.Handler) (http.Handler, error) {
	appYaml, err := ParseAppYaml(appYamlPath)
	if err != nil {
		return nil, err
	}

	handlers := appYaml.Handlers.Select(func(i *AppYamlHandler) bool {
		// StaticDir is not supported yet.
		return i.StaticFiles != ""
	})

	if err := handlers.Each(func(i *AppYamlHandler) error {
		return i.Setup()
	}); err != nil {
		return nil, err
	}

	return handlers.NewHandler(defaultHandler), nil
}
