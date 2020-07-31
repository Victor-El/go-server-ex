package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// We have to create a fileDir obj and pass to file server
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle is used for file server while REST uses HandleFunc
	http.Handle("/", fileServer)

	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/hello", welcomeHandler)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not supported", http.StatusNotFound)
			return			
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error with form submission", http.StatusNotFound)
			return			
		}

		var name string = r.FormValue("name")
		fmt.Fprintf(w, "Submitted value: " + name)
	})

	fmt.Println("Server running at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome" {
		http.Error(w, "404, not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "<h1>Welcome</h1>")
}