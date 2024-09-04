package main

import (
	"go_practice/src/15_rest_api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import ("encoding/json" "log" "net/http" "math/rand", "strconv",)

func main() {
	router := mux.NewRouter()

	routes.BookRoute(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
