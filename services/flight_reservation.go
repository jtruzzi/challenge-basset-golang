package services

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../models"
)

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
