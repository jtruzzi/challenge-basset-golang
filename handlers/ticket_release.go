package handlers

import (
	"net/http"
	"strconv"

	"../models"
	"../services"
	"github.com/julienschmidt/httprouter"
	"log"
)

// CreateTicketRelease: Endpoint for releasing product tickets to the user by email
func CreateTicketRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client, err := models.GetClient(r.Header.Get("x-client-id"))
	apiKey := r.Header.Get("X-Api-Key")
	resend, _ := strconv.ParseBool(r.URL.Query().Get("resend"))

	reservation, err := services.GetReservationWithFlightReservations(ps.ByName("reservationId"), apiKey, client.ClientId)

	if err != nil {
		NewAPIError(&APIError{false, err.Error(), http.StatusInternalServerError}, w)
		return
	}

	if len(reservation.Products) == 0 {
		NewAPIError(&APIError{false, "Product not found", http.StatusInternalServerError}, w)
		return
	}

	response, err := services.SendEmailConfirmation(reservation, resend, client)
	if err != nil {
		log.Println("[API ERROR]: SendEmailConfirmation response: ", response)
		NewAPIError(&APIError{false, "Problem sending email", http.StatusInternalServerError}, w)
		return
	}
	NewAPIResponse(&APIResponse{true, "Tickets released"}, w, http.StatusOK)
}
