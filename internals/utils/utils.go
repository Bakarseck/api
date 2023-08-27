package utils

import (
	"github.com/Bakarseck/api/internals/models"
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// The function `RenderPage` renders a web page using a template file and provided data.
func RenderPage(pagePath string, data interface{}, w http.ResponseWriter) {
	files := []string{"templates/base.html", "templates/" + pagePath + ".html"}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("🚨 " + err.Error())
	} else {
		tpl.Execute(w, data)
	}
}

// The function RenderErrorPage renders an error page with a specific error message based on the
// provided error code.
func RenderErrorPage(code int, w http.ResponseWriter) {
	errorMessage := models.Error[code]
	if errorMessage == "" {
		errorMessage = "Unknown Error"
	}

	errorData := struct {
		Message string
	}{
		Message: errorMessage,
	}

	RenderPage("error", errorData, w)
}

// The function `LoadEnv` reads an environment file, splits each line into key-value pairs, and sets
// them as environment variables.
func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
	return scanner.Err()
}

// The OpenBrowser function opens a URL in the default web browser based on the user's operating
// system.
func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return nil // Unsupported platform
	}

	return cmd.Start()
}

// The function `ValidateRequest` checks if the request is valid based on the provided URL and method,
// and returns true if it is valid.
func ValidateRequest(r *http.Request, w http.ResponseWriter, url, method string) bool {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		RenderErrorPage(http.StatusBadRequest, w)
		log.Println("400 ❌ - Bad Request")
		return false
	}
	if r.URL.Path != url {
		w.WriteHeader(http.StatusNotFound)
		RenderErrorPage(http.StatusNotFound, w)
		log.Println("404 ❌ - Page not found ", r.URL)
		return false
	}
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		log.Printf("405 ❌ - Method not allowed %s - %s on URL : %s", r.Method, method, url)
		return false
	}
	return true
}
