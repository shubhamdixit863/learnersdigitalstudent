package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// LogMiddleware logs request details
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logEntry := map[string]interface{}{
			"time":    time.Now().Format(time.RFC3339),
			"method":  r.Method,
			"path":    r.URL.Path,
			"agent":   r.UserAgent(),
			"headers": r.Header,
		}
		logData, _ := json.MarshalIndent(logEntry, "", "  ")
		log.Println(string(logData))
		next.ServeHTTP(w, r)
	})
}
