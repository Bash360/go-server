package order

import (
	"encoding/json"
	"net/http"
	"rest-api/helper"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders", GetAll).Methods("GET")
	router.HandleFunc("/orders/{id}", GetOne).Methods("GET")
	router.HandleFunc("/orders", Create).Methods("POST")
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	orders, err := getAll()
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())
	}
	helper.WithJSON(w, http.StatusOK, orders)
}

func GetOne(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if ok {
		helper.WithError(w, http.StatusBadRequest, "bad request no Id")
	}

	id, _ := strconv.Atoi(idStr)

	order, err := getOne(id)

	if err != nil {
		helper.WithError(w, http.StatusNotFound, err.Error())
	}
	helper.WithJSON(w, http.StatusOK, order)

}

func Create(w http.ResponseWriter, req *http.Request) {
	var order Order
	json.NewDecoder(req.Body).Decode(&order)
	create(&order)
}
