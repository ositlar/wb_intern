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
		http.Error(w, "gache.Get error", http.StatusInternalServerError)
		log.Fatal(err)
	}
	jsonView, jsonErr := json.Marshal(info)
	if jsonErr != nil {
		http.Error(w, "Serialize error", http.StatusInternalServerError)
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
