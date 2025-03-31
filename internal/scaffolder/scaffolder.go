package scaffolder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Directories to be created
var directories = []string{
	"services", "config", "controllers", "models",
	"routes", "middlewares",
}

// Files and their templates
var files = map[string]string{
	"main.go": "templates/main.go.tpl",
}

// CreateProject initializes the project structure inside the given project path
func CreateProject(projectPath string) {

	// Ensure the project folder exists
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		fmt.Printf("❌ Error: Project folder %s does not exist.\n", projectPath)
		return
	}

	// Create directories inside the project folder
	for _, dir := range directories {
		path := filepath.Join(projectPath, dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Printf("❌ Error creating %s: %v\n", path, err)
		} else {
			fmt.Printf("📂 Created: %s\n", path)
		}
	}

	// Generate files from templates
	for filePath, templatePath := range files {
		fullPath := filepath.Join(projectPath, filePath)
		createFileFromTemplate(fullPath, templatePath, projectPath)
	}

	fmt.Println("✅ Project successfully generated!")
}

// createFileFromTemplate processes template files
func createFileFromTemplate(filePath, templatePath, projectPath string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("❌ Error loading template %s: %v\n", templatePath, err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("❌ Error creating %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, map[string]string{"ProjectName": filepath.Base(projectPath)})
	if err != nil {
		fmt.Printf("❌ Error writing to %s: %v\n", filePath, err)
	}
}
