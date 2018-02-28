package handlers

import (
	"net/http"
	"../models"
	"../services"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func CreateTicketRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reservation, err := services.GetReservation(ps.ByName("reservationId") )
	if err != nil { panic(err) }

	var flightReservations []models.FlightReservation

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" { break }

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil { panic(err); return }

		if flightReservation.HasIssuedTicket() {
			flightReservations  = append(flightReservations, flightReservation)
		}
	}

	if len(flightReservations) == 0 { return }

	services.SendEmailConfirmation(reservation, flightReservations)
	json.NewEncoder(w).Encode(reservation.Products[0].FlightReservation)

}

