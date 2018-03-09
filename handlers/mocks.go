package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"os"
	"../models"
)

var pwd, _ = os.Getwd()

func MockedReservationsEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mockedData, _ := ioutil.ReadFile(pwd + "/mocks/reservation.json")
	var raw models.Reservation
	json.Unmarshal(mockedData, &raw)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(raw)
}

func MockedFlightReservationsEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mockedData, _ := ioutil.ReadFile(pwd + "/mocks/flight_reservation.json")
	var raw models.FlightReservation
	json.Unmarshal(mockedData, &raw)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(raw)
}
