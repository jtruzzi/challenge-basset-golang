package routes

import (
	"net/http"
	"../services"
	"github.com/gorilla/mux"
)


func ReservationRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reservationId := params["reservation_id"]

	reservation, err := services.GetReservation(reservationId )
	if err != nil { panic(err) }

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" { break }

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil {
			panic(err)
			return
		}
		product.FlightReservation = flightReservation
	}
}

