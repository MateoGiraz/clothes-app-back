package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matoegiraz/clothes-app/db"
	"github.com/matoegiraz/clothes-app/db/querys"
	"github.com/matoegiraz/clothes-app/models"
)

func handleOutfitRouting(r *mux.Router) {
	r.HandleFunc("", CreateOutfit).Methods("POST")
	r.HandleFunc("/{id}", DeleteOutfit).Methods("DELETE")
	r.HandleFunc("", GetOutfitsHandler).Methods("GET")
	r.HandleFunc("/{id}", GetOutfitHandler).Methods("GET")
}

func GetOutfitsHandler(w http.ResponseWriter, r *http.Request) {
	var outfits []models.Outfit

	rows, err := db.DB.Query(querys.GetOutfits)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.WriteHeader(http.StatusNotFound)

		return
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Outfit

		err := rows.Scan(
			&o.Id, 
			&o.TopId, 
			&o.PantsId,
			&o.ShoesId,
		)

		if err != nil {
			w.Write([]byte(fmt.Sprintf("err: %s", err)))
			w.WriteHeader(http.StatusNotFound)

			return
		}

		outfits = append(outfits, o)
	}

	if len(outfits) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There are no outfits"))

		return
	}

	json.NewEncoder(w).Encode(&outfits)
}

func GetOutfitHandler(w http.ResponseWriter, r *http.Request) {
	var o models.Outfit
	params := mux.Vars(r)
	row := db.DB.QueryRow(querys.GetOutfit, params["id"])

	err := row.Scan(
		&o.Id, 
		&o.TopId, 
		&o.PantsId, 
		&o.ShoesId,
	)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Outfit not found"))

		return
	}

	json.NewEncoder(w).Encode(&o)
}

func isValidClothingId(id int8, cat string) (err error){
	var c models.Clothing
	row := db.DB.QueryRow(querys.GetClothing, id)

	err = row.Scan(
		&c.Id,
		&c.IsAvailable, 
		&c.Name, 
		&c.Description, 
		&c.Color, 
		&c.ImageURL, 
		&c.Category,
	)

	if err != nil {
		return err
	}

	if c.Category != cat {
		return fmt.Errorf("category does not match with %s", cat)
	}

	return nil
}

func CreateOutfit(w http.ResponseWriter, r *http.Request) {
	var o models.Outfit

	json.NewDecoder(r.Body).Decode(&o)

	var err error
	err	= isValidClothingId(o.TopId, "top")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Top id err: %s", err)))

		return
	}

	err = isValidClothingId(o.PantsId, "pants")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Pants id err: %s", err)))

		return
	}

	err = isValidClothingId(o.ShoesId, "shoes")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Shoes id err: %s", err)))

		return
	}

	row := db.DB.QueryRow(
		querys.CreateOutfit,
		o.TopId,
		o.PantsId,
		o.ShoesId,
	)

	var id int64
	err = row.Scan(&id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.Write([]byte("Could not insert"))

		return
	}

	json.NewEncoder(w).Encode(id)

}

func DeleteOutfit(w http.ResponseWriter, r *http.Request) {

	var o models.Outfit
	params := mux.Vars(r)
	row := db.DB.QueryRow(querys.GetOutfit, params["id"])

	err := row.Scan(
		&o.Id, 
		&o.TopId, 
		&o.PantsId, 
		&o.ShoesId,
	)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Outfit not found"))

		return
	}

	_, err = db.DB.Exec(querys.DeleteOutfit, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Outfit not found"))

		return
	}

	w.WriteHeader(http.StatusOK)
}
