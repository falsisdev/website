package handlers

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var tmpl *template.Template

func InitTemplates() error {
	templateRoot, err := findTemplateRoot()
	if err != nil {
		return err
	}

	var files []string

	patterns := []string{
		filepath.Join(templateRoot, "templates", "*.html"),
		filepath.Join(templateRoot, "templates", "layouts", "*.html"),
		filepath.Join(templateRoot, "templates", "components", "*.html"),
	}

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}
		files = append(files, matches...)
	}

	if len(files) == 0 {
		return fmt.Errorf("No HTMl template files found")
	}

	tmpl, err = template.New("").ParseFiles(files...)
	return err
}

func findTemplateRoot() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for _, candidate := range []string{
		filepath.Join(workingDir, "web"),
		filepath.Join(workingDir, "..", "..", "web"),
	} {
		if _, err := os.Stat(filepath.Join(candidate, "templates")); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("web/templates directory not found from %s", workingDir)
}
