package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"yemeksepeti-keyvalue/internal/model"
	"yemeksepeti-keyvalue/internal/service"
)

type StorageAPI struct {
	StorageService service.StorageService
}

func NewStorageAPI(s service.StorageService) StorageAPI {
	return StorageAPI{StorageService: s}
}

/*
saves the data in memory if there is a recorded file. afterward,
it saves the data in memory to a file every 1 minute.
 */
func (s StorageAPI) FileOperations() {
	s.StorageService.FileOps()
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				s.StorageService.SaveMemoryToFile()
				log.Printf("Saved from memory to file. Registration time:", t.String())
			}
		}
	}()
}

/*
returns all data in memory as a response. only works with the GET method, it indicates this if memory is empty.
 */
func (s StorageAPI) GetAll(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)
	if req.Method != http.MethodGet {
		e := errors.New("Wrong Request. The request must be GET.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	allPairs := s.StorageService.GetAll()
	if len(allPairs) <= 0 {
		var respond model.RespondStatus
		msg := fmt.Sprintf("There is no pair.")
		respond.Status = msg
		log.Printf(msg)
		log.Printf("Status Code: %d", 200)
		JSON(w, http.StatusOK, respond)
	} else {
		log.Printf("Respond Status Code: %d", 200)
		log.Printf("Respond Data: %s", allPairs)
		JSON(w, http.StatusOK, allPairs)
	}

}
/*
Key and value pair are saved. If the relevant key is used, it is not saved and this is stated in the response.
It only works with the POST method and takes JSON data.
 */
func (s StorageAPI) Set(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)
	if req.Method != http.MethodPost {
		e := errors.New("Wrong Request. The request must be POST.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	var respondMap model.RespondMap
	body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(body, &respondMap)

	if err != nil {
		e := errors.New("There is an error in the requested data. Check the data. Data should be JSON.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	for i := 0; i < len(respondMap.Pair); i++ {
		if s.StorageService.SetIfKeyNotUsedBefore(respondMap.Pair[i].Key, respondMap.Pair[i].Value) {
			e := errors.New("Key updated.")
			respondMap.Pair[i].Status = e.Error()
			log.Printf(respondMap.Pair[i].Key + " key is used already.")
			log.Printf("Error Message:%s", e.Error())
			continue
		}
		respondMap.Pair[i].Status = "Saved."
		log.Printf("Keys are saved.")
	}
	log.Printf("Respond Status Code: %d", 201)
	log.Printf("Respond Data: %s", respondMap)
	JSON(w, http.StatusCreated, respondMap)
}
/*
The value corresponding to the key information sent in the URL is returned in the response.
It works with the GET method.
example reqyest : endpoint-> /api/get?key=foo (GET Request)
 */
func (s StorageAPI) Get(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)
	if req.Method != http.MethodGet {
		e := errors.New("Wrong Request. The request must be GET.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	key := req.URL.Query().Get("key")
	val := s.StorageService.GetValue(key)

	var respond model.RespondData
	if key == "" || val == "" {
		e := errors.New("Bad Request. The URL may be an incorrect or there may not be a value for the key value.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
	} else {
		respond.Key = key
		respond.Value = val
		respond.Status = "OK"
		log.Printf("Respond Status Code: %d", 200)
		log.Printf("Respond Data: %s", respond)
		JSON(w, http.StatusOK, respond)
	}

}
/*
deletes all data in memory. GET request should be sent to the relevant endpoint.
 */
func (s StorageAPI) Flush(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)
	if req.Method != http.MethodGet {
		e := errors.New("Wrong Request. The request must be GET.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	s.StorageService.Flush()
	var respond model.RespondStatus
	msg := fmt.Sprintf("All datas are deleted")
	respond.Status = msg
	log.Printf("Respond Status Code: %d", 200)
	log.Printf("Respond Data: %s", respond)
	JSON(w, http.StatusOK, respond)
}
/*
Key information is obtained and the relevant key and value pair is deleted.
example: delete?key=foo
*/
func (s StorageAPI) Delete(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)
	if req.Method != http.MethodGet {
		e := errors.New("Wrong Request. The request must be GET.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}
	key := req.URL.Query().Get("key")
	if key == "" {
		e := errors.New("Bad Request. The URL may be an incorrect or there may not be a value for the key value.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
	} else {
		s.StorageService.Delete(key)
		var respond model.RespondStatus
		msg := fmt.Sprintf("The key is deleted")
		respond.Status = msg
		log.Printf("Respond Status Code: %d", 200)
		log.Printf("Respond Data: %s", respond)
		JSON(w, http.StatusCreated, respond)
	}

}
