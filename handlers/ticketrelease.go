package handlers

import (
	"net/http"
	"../services"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func CreateTicketRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

