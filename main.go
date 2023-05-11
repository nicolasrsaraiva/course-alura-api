package main

import (
	"net/http"

	"github.com/nicolasrsaraiva/course-alura-api/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
