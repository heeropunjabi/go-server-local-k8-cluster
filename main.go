package main

// Import Packages
import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Server the Desired HTML File
	http.Handle("/", http.FileServer(http.Dir("./content")))

	fmt.Println("Server is running on port 9091")

	log.Fatal(http.ListenAndServe(":9091", nil))
}
