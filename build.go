package localgaesupport

import (
	"log"
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

	log.Printf("Starting loalgaesuuport with %s\n", appYamlPath)
	for idx, i := range handlers {
		log.Printf("%s[%d]: %s", appYamlPath, idx, i.URL)
	}

	return handlers.NewHandler(defaultHandler), nil
}
