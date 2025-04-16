package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

// User file access map
var userFiles = map[string][]string{
	"alice": {"file1.txt", "file2.txt"},
	"bob":   {"file3.txt"},
}

// FileListHandler returns the list of files a user can access
func FileListHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	files, ok := userFiles[user]
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	jsonResponse, _ := json.Marshal(files)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// FileDataHandler returns file content if the user has access
func FileDataHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	fileName := r.URL.Query().Get("file")

	allowedFiles, ok := userFiles[user]
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !contains(allowedFiles, fileName) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	filePath := filepath.Join("files", fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

// Helper function to check if a slice contains an element
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
