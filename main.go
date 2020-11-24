package main

import (
	"fmt"
	"net/http"

	"github.com/wanderporto/api-json-crud/routes"
)

func main() {
	fmt.Println("Listening port 3000")

	r := routes.NewRouter()
	http.ListenAndServe(":3000", r)
}
