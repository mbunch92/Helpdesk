package main

import (
	"fmt"
	"net/http"
)

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createAccount" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		// Serve the create account form
		fmt.Fprint(w, "Create Account Form")
	case "POST":
		// Handle the form submission
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		fmt.Println("Username:", username)
		fmt.Println("Password:", password)
		fmt.Println("Email:", email)

		// Display success message
		fmt.Fprintln(w, "Account created successfully!")
	default:
		// Invalid request method
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
