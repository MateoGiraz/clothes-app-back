package routes

import(
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matoegiraz/clothes-app/models"
	"github.com/matoegiraz/clothes-app/startup"
)

func handleClothingTypeRouting(r *mux.Router) {
	r.HandleFunc("", GetClothingType).Methods("GET")

}

func GetClothingType(w http.ResponseWriter, r *http.Request) {
	var clothingTypeArr []models.AvailableClothesType
	clothingTypeArr = startup.CreateClothingType()
	
	json.NewEncoder(w).Encode(&clothingTypeArr)
}
