package routes

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	c := r.PathPrefix("/clothes").Subrouter()
	o := r.PathPrefix("/outfits").Subrouter()

	handleClothesRouting(c)
	handleOutfitRouting(o)

	return r
}
