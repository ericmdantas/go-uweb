package uweb

import "strings"

func normalizePath(path string) string {
	normalizedPath := strings.Replace(strings.ToLower(path), " ", "", -1)

	if len(normalizedPath) == 0 {
		normalizedPath = "/"
	}

	if len(normalizedPath) > 1 && normalizedPath[len(normalizedPath)-1:] == "/" {
		normalizedPath = normalizedPath[:len(normalizedPath)-1]
	}

	if normalizedPath[:1] != "/" {
		normalizedPath = "/" + normalizedPath[:len(normalizedPath)]
	}

	return normalizedPath
}
