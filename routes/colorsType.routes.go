package routes

import(
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matoegiraz/clothes-app/models"
	"github.com/matoegiraz/clothes-app/startup"
)

func handleColorTypeRouting(r *mux.Router) {
	r.HandleFunc("", GetColorsType).Methods("GET")

}

func GetColorsType(w http.ResponseWriter, r *http.Request) {
	var colorTypeArr []models.AvailableColorType
	colorTypeArr = startup.CreateColors()
	
	json.NewEncoder(w).Encode(&colorTypeArr)
}
