package main

import (
	"fmt"
	"net/http"
	"std/omkesh/Practice/workers/database"

	// "std/omkesh/Practice/workers/endpoints"

	endpoints "std/omkesh/Practice/workers/endpoints"
	handler "std/omkesh/Practice/workers/handlers"
	repository "std/omkesh/Practice/workers/repositories"
	service "std/omkesh/Practice/workers/service"

	"github.com/gorilla/mux"
)

const (
	HTTPAddr = "9000"
)

func main() {

	router := mux.NewRouter()

	//MysqlDB Dependancy
	conPool := database.NewMysqlConnection()

	taskRepository := repository.NewRepositoryImpl(conPool)
	taskService := service.NewServiceImpl(taskRepository)
	taskHandler := handler.NewHandlerImpl(taskService)
	endpoints.NewRoute(router, taskHandler)
	fmt.Println(http.ListenAndServe(":"+HTTPAddr, router))
}
