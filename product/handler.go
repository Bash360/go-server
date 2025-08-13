package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"rest-api/helper"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", GetAll).Methods("GET")
	router.HandleFunc("/products/{id}", GetOne).Methods("GET")
	router.HandleFunc("/products", Create).Methods("POST")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	helper.WithJSON(w, http.StatusOK, getAll())
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		log.Println("Id does not exist")
		helper.WithError(w, http.StatusNotFound, "Id does not exist")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		helper.WithError(w, http.StatusBadRequest, err.Error())
	}
	product, err := getOne(intId)
	if err != nil {
		helper.WithError(w, http.StatusNotFound, err.Error())
	}
	helper.WithJSON(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)

	newProduct, err := create(&product)
	if err != nil {
		log.Println(err.Error())

		helper.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WithJSON(w, http.StatusCreated, newProduct)
}
