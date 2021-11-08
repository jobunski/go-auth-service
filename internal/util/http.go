package util

import "strings"

func ReturnsJson(path, accept string) bool {
	return strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/token")

}
