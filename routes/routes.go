package routes

import (
	"github.com/julienschmidt/httprouter"
	"../handlers"
)

func BuildRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/reservations/:reservationId/ticket-release", handlers.CreateTicketRelease)
	MockedServices(router)

	return router
}

func MockedServices(router *httprouter.Router) {
	router.GET("/reservations/:reservationId", handlers.MockedReservationsEndpoint)
	router.GET("/flights/reservations/:reservationId", handlers.MockedFlightReservationsEndpoint)
}
