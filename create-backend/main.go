package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var structure = []string{
	"cmd/main.go",
	"config/config.go",
	"config/config.yaml",
	"internal/auth/auth.go",
	"internal/auth/middleware.go",
	"internal/database/database.go",
	"internal/database/migrations/.gitkeep",
	"internal/models/user.go",
	"internal/models/project.go",
	"internal/repositories/user_repo.go",
	"internal/repositories/project_repo.go",
	"internal/services/user_service.go",
	"internal/services/project_service.go",
	"internal/handlers/user_handler.go",
	"internal/handlers/project_handler.go",
	"internal/middleware/logging.go",
	"internal/middleware/cors.go",
	"pkg/utils/hash.go",
	"pkg/utils/jwt.go",
	"pkg/responses/response.go",
	"api/routes.go",
	"docs/api-docs.md",
	"tests/user_test.go",
	"tests/project_test.go",
	"go.mod",
	"go.sum",
	"Makefile",
	"README.md",
}

func createFilesAndDirs() {
	for _, path := range structure {
		fullPath := filepath.Join(".", path)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}

		file, err := os.Create(fullPath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", fullPath, err)
			continue
		}
		defer file.Close()

		if err := writeBoilerplate(fullPath, file); err != nil {
			fmt.Printf("Error writing to file %s: %v\n", fullPath, err)
		}
	}
}

func writeBoilerplate(path string, file *os.File) error {
	content := ""
	switch filepath.Base(path) {
	case "main.go":
		content = `package main

import "fmt"

func main() {
	fmt.Println("BaaS Backend Started")
}`
	case "config.go":
		content = `package config

import "fmt"

func LoadConfig() {
	fmt.Println("Loading Config")
}`
	case "Makefile":
		content = `run:
	go run cmd/main.go`
	case "README.md":
		content = `# Golang BaaS Backend

Generated by CLI.`
	}

	_, err := file.WriteString(content)
	return err
}

func main() {
	fmt.Println("Creating Golang BaaS backend structure...")
	createFilesAndDirs()
	fmt.Println("Structure created successfully!")
}
