package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Request body struct
type ValidationRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Response struct
type ValidationResponse struct {
	Success bool              `json:"success"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Validation functions
func validateName(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return "Name cannot be empty"
	}
	if len(name) < 2 {
		return "Name must have at least 2 characters"
	}
	return ""
}

func validateEmail(email string) string {
	// Basic email regex
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return "Invalid email format"
	}
	return ""
}

func validatePhone(phone string) string {
	// Accept only digits, 10–15 length
	re := regexp.MustCompile(`^[0-9]{10,15}$`)
	if !re.MatchString(phone) {
		return "Phone number must be 10–15 digits"
	}
	return ""
}

// Handler function
func validationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ValidationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	errors := make(map[string]string)

	if msg := validateName(req.Name); msg != "" {
		errors["name"] = msg
	}
	if msg := validateEmail(req.Email); msg != "" {
		errors["email"] = msg
	}
	if msg := validatePhone(req.Phone); msg != "" {
		errors["phone"] = msg
	}

	res := ValidationResponse{
		Success: len(errors) == 0,
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	if res.Success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/api/validations", validationHandler)

	fmt.Println("🚀 Server running on http://localhost:3000/api/validations")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
