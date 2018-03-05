package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"../services"
	"github.com/julienschmidt/httprouter"
)

func CreateTicketRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reservation, err := services.GetReservation(ps.ByName("reservationId"))
	if err != nil {
		panic(err)
	}

	resend, _ := strconv.ParseBool(r.URL.Query().Get("resend"))

	var products []models.Product

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" {
			break
		}

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil {
			panic(err)
			return
		}

		if flightReservation.HasIssuedTicket() {
			product.FlightReservation = flightReservation
			products = append(products, product)
		}
	}

	if len(products) == 0 {
		return
	}

	reponse, err := services.SendEmailConfirmation(reservation, products, resend)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(reponse)
}
