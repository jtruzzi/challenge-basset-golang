package routes

import (
	"net/http"
	//"github.com/gorilla/mux"
	"../services"
	"encoding/json"
	"fmt"
)


// MovieRoute
func ReservationRoute(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	reservation, err := services.GetReservation()
	if err != nil {
		panic(err)
	}
	jsonReservation, _ := json.MarshalIndent(reservation, "", "  ")
	//json.NewEncoder(w).Encode(jsonReservation)
	fmt.Fprintf(w, string(jsonReservation))
	for _, product := range reservation.Products {
		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil {
			json.NewEncoder(w).Encode(err2)
		}

		//jsonValue, _ := json.MarshalIndent(flightReservations, "", "  ")
		//fmt.Fprintf(w, string(jsonValue))
		jsonFlightReservation, _ := json.MarshalIndent(flightReservation, "", "  ")
		//json.NewEncoder(w).Encode(jsonFlightReservation)
		fmt.Fprintf(w, string(jsonFlightReservation))
	}
}

