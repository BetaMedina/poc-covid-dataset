package main

import (
	"dataset-covid/pkg/api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routes.GenerateMuxRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), r))
}
