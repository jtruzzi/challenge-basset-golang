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

	var products []models.Product

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" { break }

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil { panic(err); return }

		if flightReservation.HasIssuedTicket() {
			product.FlightReservation = flightReservation
			products  = append(products, product)
		}
	}

	if len(products) == 0 { return }

	services.SendEmailConfirmation(reservation, products)
	json.NewEncoder(w).Encode(reservation.Products[0].FlightReservation)

}

