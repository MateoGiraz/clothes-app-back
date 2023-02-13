package main

import (
	"net/http"

	"github.com/matoegiraz/clothes-app/db"
	"github.com/matoegiraz/clothes-app/routes"
)

func main() {
	r := routes.GetRouter()

	db.Connect()

	http.ListenAndServe(":2000", r)
}
