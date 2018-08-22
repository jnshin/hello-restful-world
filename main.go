// main
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello RESTful World!")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Printf("OK. Port %s will be used.\n", port)
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q",
				html.EscapeString(r.URL.Path))
		})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
