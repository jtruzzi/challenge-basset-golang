package services

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../models"
	"log"
)

func GetReservation(reservationId string) (models.Reservation, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.basset.ws/reservations/"+reservationId+"?site=AR&channel=WEB", nil)

	request.Header.Add("X-Api-Key", os.Getenv("X_API_KEY"))
	request.Header.Add("x-client-id", os.Getenv("X_CLIENT_ID"))

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

func GetFlightReservation(reservationId string) (models.FlightReservation, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.basset.ws/flights/reservations/"+reservationId+"?site=AR&channel=WEB", nil)

	request.Header.Add("X-Api-Key", os.Getenv("X_API_KEY"))
	request.Header.Add("x-client-id", os.Getenv("X_CLIENT_ID"))

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
