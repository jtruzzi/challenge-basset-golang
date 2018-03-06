package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"../services"
	"github.com/julienschmidt/httprouter"
)

// CreateTicketRelease: Endpoint for releasing product tickets to the user by email
func CreateTicketRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client, err := models.GetClient(r.Header.Get("x-client-id"))
	apiKey := r.Header.Get("X-Api-Key")
	resend, _ := strconv.ParseBool(r.URL.Query().Get("resend"))

	reservation, err := services.GetReservationWithFlightReservations(ps.ByName("reservationId"), apiKey, client.ClientId)

	if err != nil {
		return
	}

	if len(reservation.Products) == 0 {
		return
	}

	response, err := services.SendEmailConfirmation(reservation, resend, client)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}
