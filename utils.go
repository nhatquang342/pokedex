package main

import "strings"

func normalizeURL(url string) string {
	return strings.TrimSuffix(url, "/")
}