// web app is running in a directory which has a subdirectory named 'public/' and

package main

import (
	"fmt"
	"net/http"
)

var webPort string = ":8080"

func main() {
	fmt.Println("Go Web Server running.")
	fmt.Println("Listening on port " + webPort)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(webPort, nil)
}
