package actions

import (
	"net/http"
	"../services"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)


func EmailConfirmation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reservation, err := services.GetReservation(ps.ByName("reservationId") )
	if err != nil { panic(err) }

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" { break }

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil { panic(err); return }

		services.SendEmailConfirmation(reservation, product, flightReservation)
		json.NewEncoder(w).Encode(reservation.Products[0].FlightReservation)
	}
}

