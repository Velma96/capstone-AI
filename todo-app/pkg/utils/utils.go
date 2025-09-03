package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ParseDate parses a date string in YYYY-MM-DD format
func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", dateStr)
}

// FormatDate formats a time.Time to YYYY-MM-DD string
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

// ReadJSON reads and decodes JSON from request body
func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return fmt.Errorf("request body must only contain a single JSON object")
	}

	return nil
}

// WriteJSON writes JSON response
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	return err
}

// ErrorJSON writes error response in JSON format
func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusInternalServerError
	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}{
		Error:   true,
		Message: err.Error(),
	}

	WriteJSON(w, statusCode, payload)
}

// IsValidStatus validates if a status string is valid
func IsValidStatus(status string) bool {
	switch status {
	case "pending", "in_progress", "completed":
		return true
	default:
		return false
	}
}