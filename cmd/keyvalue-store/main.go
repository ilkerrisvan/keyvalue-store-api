package main

import (
	"fmt"
	"log"
	"net/http"
	"yemeksepeti-keyvalue/internal/api"
	"yemeksepeti-keyvalue/internal/repository"
	"yemeksepeti-keyvalue/internal/service"
)

/*
App runs on localhost, port 8000.
 */
func main() {
	e := run(8000)
	if e != nil {
		log.Printf("Connection failed.")
	}
}

/*
it first checks if there is a saved file. if there is a recorded file, it saves it to memory.
 */
func run(port int) error {
	storageAPI := InitStorageAPI()
	//saves to memory if exist file in temp/...-.json
	storageAPI.FileOperations()
	log.Printf("Server running at http://localhost:%d/", port)
	http.HandleFunc("/api/set", storageAPI.Set)
	http.HandleFunc("/api/get", storageAPI.Get)
	http.HandleFunc("/api/get-all", storageAPI.GetAll)
	http.HandleFunc("/api/flush", storageAPI.Flush)
	http.HandleFunc("/api/delete", storageAPI.Delete)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}
	return nil
}

/*
the application works with different layers. when the API in the top layer is initialized, all layers work.
 */
func InitStorageAPI() api.StorageAPI {
	storageRepository := repository.NewRepository()
	storageService := service.NewConverterService(storageRepository)
	storageAPI := api.NewStorageAPI(storageService)
	return storageAPI
}
