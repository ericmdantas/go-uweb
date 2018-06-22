package uweb

import "strings"

func normalizePath(path string) string {
	pathNormalized := strings.Replace(strings.ToLower(path), " ", "", -1)

	if len(pathNormalized) == 0 {
		pathNormalized = "/"
	}

	if len(pathNormalized) > 1 && pathNormalized[len(pathNormalized)-1:] == "/" {
		pathNormalized = pathNormalized[:len(pathNormalized)-1]
	}

	if pathNormalized[:1] != "/" {
		pathNormalized = "/" + pathNormalized[:len(pathNormalized)]
	}

	return pathNormalized
}
