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

func handleClothesRouting(r *mux.Router) {
	r.HandleFunc("", CreateClothing).Methods("POST")
	r.HandleFunc("/{id}", DeleteClothing).Methods("DELETE")
	r.HandleFunc("", GetClothesHandler).Methods("GET")
	r.HandleFunc("/{id}", GetClothingHandler).Methods("GET")
}

func GetClothesHandler(w http.ResponseWriter, r *http.Request) {
	var clothes []models.Clothing

	rows, err := db.DB.Query(querys.GetClothes)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.WriteHeader(http.StatusNotFound)

		return
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Clothing

		err := rows.Scan(&c.Id, &c.IsAvailable, &c.Name, &c.Description, &c.Color, &c.Category)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("err: %s", err)))
			w.WriteHeader(http.StatusNotFound)

			return
		}

		clothes = append(clothes, c)
	}

	if len(clothes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There are no clothes"))

		return
	}

	json.NewEncoder(w).Encode(&clothes)
}

func GetClothingHandler(w http.ResponseWriter, r *http.Request) {
	var c models.Clothing
	params := mux.Vars(r)
	row := db.DB.QueryRow(querys.GetClothing, params["id"])

	err := row.Scan(&c.Id, &c.IsAvailable, &c.Name, &c.Description, &c.Color, &c.Category)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Clothing not found"))

		return
	}

	json.NewEncoder(w).Encode(&c)
}

func CreateClothing(w http.ResponseWriter, r *http.Request) {
	var c models.Clothing

	json.NewDecoder(r.Body).Decode(&c)

	row := db.DB.QueryRow(
		querys.CreateClothing,
		c.IsAvailable,
		c.Name,
		c.Description,
		c.Color,
		c.Category,
	)

	var id int64
	err := row.Scan(&id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.Write([]byte("Could not insert"))

		return
	}

	json.NewEncoder(w).Encode(id)

}

func DeleteClothing(w http.ResponseWriter, r *http.Request) {

	var c models.Clothing
	params := mux.Vars(r)
	row := db.DB.QueryRow(querys.GetClothing, params["id"])

	err := row.Scan(&c.Id, &c.IsAvailable, &c.Name, &c.Description, &c.Color, &c.Category)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Clothing not found"))

		return
	}

	_, err = db.DB.Exec(querys.DeleteClothing, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Clothing not found"))

		return
	}

	w.WriteHeader(http.StatusOK)
}
