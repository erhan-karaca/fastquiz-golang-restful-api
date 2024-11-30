package utils

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func GetRootPath() string {
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Failed to get root path: %v", err)
	}
	return rootPath
}

func FormatSlug(name string, suffix string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")
	return slug + "-" + suffix
}
