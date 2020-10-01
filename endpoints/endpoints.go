package endpoints

import (
	handler "practice/whetherinfo/workers/handlers"

	"github.com/gorilla/mux"
)

//NewRoute All Application Routes Are defiend Here
func NewRoute(router *mux.Router, handler *handler.HandlersImpl) {
	router.HandleFunc("/city", handler.CreateCities).Methods("POST")
	router.HandleFunc("/city/{cityName}", handler.GetCity).Methods("GET")
}
