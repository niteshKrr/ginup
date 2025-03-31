package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// BasicSetup initializes a new Gin project with the given name
func BasicSetup(projectName string) (string, error) {
	projectPath, err := filepath.Abs(projectName)
	if err != nil {
		return "", fmt.Errorf("❌ Error getting absolute path: %v", err)
	}

	fmt.Println("🔥 Creating Gin project:", projectName)

	// Create project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return "", fmt.Errorf("❌ Error creating project directory: %v", err)
	}

	// Initialize Go module inside the project directory
	if err := runCommand(projectPath, "go", "mod", "init", projectName); err != nil {
		return "", fmt.Errorf("❌ Error initializing Go module: %v", err)
	}

	done := make(chan bool)

	// Start a spinner animation
	go spinner(done)

	// Install Gin
	err = runCommand(projectPath, "go", "get", "-u", "github.com/gin-gonic/gin")

	// Stop the spinner
	done <- true
	// time.Sleep(200 * time.Millisecond)

	if err != nil {
		return "", fmt.Errorf("❌ Error installing Gin: %v", err)
	}
	
	return projectPath, nil
}

// runCommand executes a command inside the specified directory
func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir // Ensure the command runs in the correct directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func spinner(done chan bool) {
	chars := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")  // Move to the beginning of the line
			fmt.Println("📦 Installing Gin... Done!") // Ensure it prints a new line
			return
		default:
			fmt.Printf("\r📦 Installing Gin %s", chars[i%len(chars)])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}
