package uweb

import "strings"

func NormalizePath(path string) string {
	path = strings.Replace(strings.ToLower(path), " ", "", -1)

	if len(path) == 0 {
		path = "/"
	}

	if len(path) > 1 && path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}

	if len(path) >= 1 && path[:1] != "/" {
		tmp := "/" + path[:len(path)]
		path = tmp
	}

	return path
}