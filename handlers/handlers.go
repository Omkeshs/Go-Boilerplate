package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	service "practice/whetherinfo/workers/service"

	"github.com/gorilla/mux"
)

//HandlersImpl for handler Functions
type HandlersImpl struct {
	svc service.Service
}

//NewHandlerImpl inits dependancies for graphQL and Handlers
func NewHandlerImpl(service service.Service) *HandlersImpl {
	return &HandlersImpl{svc: service}
}

//Create handler function
func (handlersImpl HandlersImpl) CreateCities(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var wg sync.WaitGroup
	go handlersImpl.svc.CreateCity(ctx, &wg)

	resp := "Creating cities at background."

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}

//Update handler function
func (handlersImpl HandlersImpl) GetCity(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := req.Context()

	vars := mux.Vars(req)
	cityName := vars["cityName"]
	resp, err := handlersImpl.svc.GetCity(ctx, cityName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
