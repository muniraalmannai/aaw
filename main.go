package main

import (
	"asciiart/asciiart"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	// Update the import path to match your module structure
)

var banners = map[string]string{
	"shadow":     "banners/shadow.txt",
	"standard":   "banners/standard.txt",
	"thinkertoy": "banners/thinkertoy.txt",
}

func main() {
	http.HandleFunc("/", homeHandler)              // Handle requests to the home page
	http.HandleFunc("/ascii-art", asciiArtHandler) // Handle form submissions
	http.ListenAndServe(":8080", nil)              // Start the server on port 8080
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html") // Load the HTML template
	if err != nil {                                          // If there's an error loading the template, handle it
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil) // Render the template and send it to the browser
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // Ensure the request method is POST
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")     // Get the text from the form
	banner := r.FormValue("banner") // Get the banner choice from the form

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	asciiArt, err := generateAsciiArt(text, banner) // Generate ASCII art
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		AsciiArt string
		Text     string
		Banner   string
	}{
		AsciiArt: asciiArt, // The generated ASCII art
		Text:     text,
		Banner:   banner,
	}

	tmpl.Execute(w, data) // Render the template with the ASCII art
}

func generateAsciiArt(text, banner string) (string, error) {
	bannerFile := filepath.Join("banners", fmt.Sprintf("%s.txt", banner)) // Determine the file path for the selected banner
	modifiedInput := ModifyString(text)                                   // Clean up the input string
	asciiArt, err := asciiart.AsciiTable(modifiedInput, bannerFile)       // Generate ASCII art
	if err != nil {
		return "", err
	}
	return asciiArt, nil
}

func ModifyString(input string) string {
	modifiedString := ""
	for _, char := range input {
		if char >= ' ' && char <= '~' {
			modifiedString += string(char)
		}
	}
	return modifiedString
}
