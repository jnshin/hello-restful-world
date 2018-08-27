// main
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	/* buffer 없는 channel 로 동기화 */
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	currTime := time.Now()

	fmt.Println("Hello RESTful World!")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Printf("OK. Port %s will be used.\n", port)
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q at %s",
				html.EscapeString(r.URL.Path), currTime.String())
		})

	/* Setup sub-routines */

	go func() {
		for {
			currTime = <-ticker.C
		}
	}()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
