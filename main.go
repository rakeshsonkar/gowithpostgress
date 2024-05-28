package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rakesh/go-postgress/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 9090")

	log.Fatal(http.ListenAndServe(":9090", r))
}
