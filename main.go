package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coderajay94/golang-jwt/handlers"
)

func main() {
	fmt.Println("welcome to the jwt authentication demo server")

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	//http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
