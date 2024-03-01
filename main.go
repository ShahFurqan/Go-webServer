// Import necessary packages
package main

// Import required libraries
import (
	"fmt"
	"log"
	"net/http"
)

// Define a function to handle form submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request
	if err := r.ParseForm(); err != nil {
		// If there's an error parsing the form, return an error message
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Respond with a success message for the POST request
	fmt.Fprintf(w, "POST request successful\n")

	// Retrieve values from the form
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Display the values in the response
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Define a function to handle requests to the "/hello" path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the URL path is "/hello"
	if r.URL.Path != "/hello" {
		// If not, return a 404 not found error
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the request method is "GET"
	if r.Method != "GET" {
		// If not, return a 404 not found error
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Respond with a simple "hello!" message for a successful GET request
	fmt.Fprintf(w, "hello!")
}

// The main function where the server configuration is set up
func main() {
	// Create a file server for the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle requests with the file server at the root path
	http.Handle("/", fileServer)

	// Register the formHandler for requests to the "/form" path
	http.HandleFunc("/form", formHandler)

	// Register the helloHandler for requests to the "/hello" path
	http.HandleFunc("/hello", helloHandler)

	// Print a message indicating that the server is starting
	fmt.Printf("starting server at port 8080\n")

	// Start the server on port 8080, log any errors
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
