package routes

import (
	"github.com/julienschmidt/httprouter"
	"../handlers"
)

func BuildRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/reservations/:reservationId/ticket-release", handlers.CreateTicketRelease)

	return router
}
