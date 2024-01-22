package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *server) HandleShowOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	info, err := s.cache.Get(id)
	if err != nil {
		http.Error(w, "cache.Get error: "+err.Error(), http.StatusInternalServerError)
	}
	jsonView, jsonErr := json.Marshal(info)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Fprint(w, string(jsonView))
}

func (s *server) HandleStart(w http.ResponseWriter, r *http.Request) {
	arr := s.cache.GetAll()
	jsonView, jsonErr := json.Marshal(arr)
	if jsonErr != nil {
		http.Error(w, "Serialize error", http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(jsonView))
}
