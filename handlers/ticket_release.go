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

	client, err := models.GetClient(r.Header.Get("x-client-id"))
	apiKey := r.Header.Get("X-Api-Key")

	reservation, err := services.GetReservation(ps.ByName("reservationId"), apiKey, client.ClientId)
	if err != nil {
		panic(err)
	}

	resend, _ := strconv.ParseBool(r.URL.Query().Get("resend"))

	var products []models.Product

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" {
			break
		}

		flightReservation, err := services.GetFlightReservation(product.ReservationId, apiKey, client.ClientId)
		if err != nil {
			panic(err)
			break
		}

		if flightReservation.HasIssuedTicket() {
			product.FlightReservation = flightReservation
			products = append(products, product)
		}
	}

	if len(products) == 0 {
		return
	}

	reponse, err := services.SendEmailConfirmation(reservation, products, resend, client)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(reponse)
}
