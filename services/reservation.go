package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../models"
)

// GetReservationWithFlightReservations: Fetches Reservation, including FlightReservation in products
func GetReservationWithFlightReservations(reservationId string, apiKey string, clientId string) (models.Reservation, error) {
	reservation, err := GetReservation(reservationId, apiKey, clientId)

	if err != nil {
		log.Println("Couldn't fetch reservation")
		return reservation, err
	}

	for index, _ := range reservation.Products {
		product := &reservation.Products[index]

		if product.Type != "FLIGHT" {
			log.Println("Not valid product type")
			break
		}

		flightReservation, err := GetFlightReservation(product.ReservationId, apiKey, clientId)
		if err != nil {
			log.Println("Couldn't fetch flight reservation")
			break
		}

		if flightReservation.HasIssuedTicket() {
			product.FlightReservation = flightReservation
		}
	}
	return reservation, nil
}

// GetReservation: Fetches Reservation
func GetReservation(reservationId string, apiKey string, clientId string) (models.Reservation, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s/%s?site=AR&channel=WEB", os.Getenv("BASSET_API"), "reservations", reservationId)

	request, err := http.NewRequest("GET", url, nil)

	request.Header = map[string][]string{
		"X-Api-Key":   {apiKey},
		"x-client-id": {clientId},
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
		return models.Reservation{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	reservation := models.Reservation{}
	errMarshall := json.Unmarshal([]byte(data), &reservation)

	if errMarshall != nil {
		log.Panic(errMarshall)
		return models.Reservation{}, errMarshall
	}

	return reservation, nil
}

// GetFlightReservation: Fetches FlightReservation
func GetFlightReservation(reservationId string, apiKey string, clientId string) (models.FlightReservation, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s/%s?site=AR&channel=WEB", os.Getenv("BASSET_API"), "flights/reservations", reservationId)

	request, err := http.NewRequest("GET", url, nil)

	request.Header = map[string][]string{
		"X-Api-Key":   {apiKey},
		"x-client-id": {clientId},
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
		return models.FlightReservation{}, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var flightReservation models.FlightReservation
	errMarshall := json.Unmarshal([]byte(data), &flightReservation)
	if errMarshall != nil {
		return models.FlightReservation{}, errMarshall
	}
	return flightReservation, nil
}
