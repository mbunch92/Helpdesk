package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmp1 := template.Must(template.ParseFiles("templates/index.html"))
	tmp1.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/createAccount", createAccountHandler)

	connStr := "user=postgres dbname=Helpdesk password=Illidan12! host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	http.ListenAndServe(":8080", nil)
}
