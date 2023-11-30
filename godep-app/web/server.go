// server.go

package main

import (
	"fmt"
	"github.com/Radha-Sundar/godep-app/pkg/greeter"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := greeter.GetGreeting()
	fmt.Fprintf(w, "<h1>%s</h1>", message)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	http.ListenAndServe(":8080", nil)
}
