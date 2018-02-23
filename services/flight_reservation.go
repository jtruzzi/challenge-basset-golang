package services

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../models"
	"os"
	"fmt"
)

func GetFlightReservation(reservation_id string) (models.FlightReservation, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.basset.ws/flights/reservations/" + reservation_id + "?site=AR&channel=WEB", nil)

	request.Header.Add("X-Api-Key", os.Getenv("X-Api-Key"))
	request.Header.Add("x-client-id", os.Getenv("x-client-id"))

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
		return models.FlightReservation{}, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	// fmt.Println(string(data))
	var flightReservations models.FlightReservation
	errMarshall := json.Unmarshal([]byte(data), &flightReservations)
	if errMarshall != nil {
		return models.FlightReservation{}, errMarshall
	}
	return flightReservations, nil
}