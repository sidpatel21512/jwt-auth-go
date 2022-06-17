package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sidpatel21512/jwt-auth-go/router"
)

func main() {
	fmt.Println("Welcome to jwt-auth-go project")

	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
