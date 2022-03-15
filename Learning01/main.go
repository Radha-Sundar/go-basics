package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {

	//var (
	//	ListenAddr = "localhost:8080"
	//	RedisAddr  = "localhost:6379"
	//)
	//
	//database, err := db.NewDatabase(RedisAddr)
	//if err != nil {
	//	log.Fatalf("Failed to connect to redis: %s", err.Error())
	//}
	//router := initRouter(database)
	//router.Run(ListenAddr)

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/health_check", check).Methods("GET")

	//http.HandleFunc("/", index)
	//http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", router)
}
