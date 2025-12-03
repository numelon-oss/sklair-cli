package discovery

import (
	"os"
	"strings"
)

func ComponentDiscovery(source string) (map[string]string, error) {
	// TODO: make the components path configurable as per sklair.json later
	dir, err := os.ReadDir(source)
	if err != nil {
		return nil, err
	}

	components := make(map[string]string)

	for _, file := range dir {
		if !file.IsDir() {
			components[strings.ToLower(file.Name())] = file.Name()
		}
	}

	return components, nil
}
