package services

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"../models"
	"os"
)

func GetReservation(reservationId  string) (models.Reservation, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.basset.ws/reservations/" + reservationId  + "?site=AR&channel=WEB", nil)

	request.Header.Add("X-Api-Key", os.Getenv("X-Api-Key"))
	request.Header.Add("x-client-id", os.Getenv("x-client-id"))

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
		return models.Reservation{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	reservation := models.Reservation{}

	errMarshall := json.Unmarshal([]byte(data), &reservation)

	if errMarshall  != nil {
		log.Panic(errMarshall)
		return models.Reservation{}, errMarshall
	}

	return reservation, nil

}