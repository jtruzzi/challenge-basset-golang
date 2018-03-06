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
