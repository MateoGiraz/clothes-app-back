package routes

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	c := r.PathPrefix("/clothing").Subrouter()
	o := r.PathPrefix("/outfits").Subrouter()

	aclt := r.PathPrefix("/availableClothingType").Subrouter()
	acot := r.PathPrefix("/availableColorsType").Subrouter()

	handleClothingRouting(c)
	handleOutfitRouting(o)
	
	handleClothingTypeRouting(aclt)
	handleColorTypeRouting(acot)
	
	return r
}
