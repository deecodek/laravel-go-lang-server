package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func main() {
	// Set Laravel public directory
	laravelPublicDir := "./laravel/public"

	// Route requests to Laravel
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveLaravel(w, r, laravelPublicDir)
	})

	// Start the server
	port := ":8000"
	fmt.Printf("Laravel app running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func serveLaravel(w http.ResponseWriter, r *http.Request, laravelPublicDir string) {
	// Determine the file path
	requestPath := r.URL.Path
	fullPath := filepath.Join(laravelPublicDir, requestPath)

	// Serve static files (like CSS, JS, images)
	if filepath.Ext(fullPath) != "" {
		http.ServeFile(w, r, fullPath)
		return
	}

	// Route all other requests to Laravel's public/index.php
	cmd := exec.Command("php", filepath.Join(laravelPublicDir, "index.php"))
	cmd.Env = append(cmd.Env, fmt.Sprintf("REQUEST_URI=%s", r.URL.RequestURI()))

	// Capture PHP output
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error processing Laravel request", http.StatusInternalServerError)
		fmt.Println("PHP Error:", err)
		return
	}

	// Write PHP output as response
	w.Header().Set("Content-Type", "text/html")
	w.Write(output)
}
