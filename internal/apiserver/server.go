package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"github.com/ositlar/wb_intern_1/internal/model"
	"github.com/ositlar/wb_intern_1/internal/store/cache"
	"github.com/ositlar/wb_intern_1/internal/store/sqlstore"
)

const (
	clusterID  = "wbCluster"
	clientID   = "server-sub"
	NatsSrvURL = "nats://127.0.0.1:4222"
)

type server struct {
	router *mux.Router
	store  sqlstore.Store
	sc     stan.Conn
	cache  cache.Cache
}

func newServer(store sqlstore.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
		cache:  *cache.NewCache(&store),
	}
	s.configureRouter()
	s.configureStan()
	go s.startListentStanServer()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/orders", s.HandleShowOrder)
	s.router.HandleFunc("/show", s.HandleStart)
}

func (s *server) configureStan() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(NatsSrvURL))
	if err != nil {
		log.Fatal(err)
	}
	s.sc = sc
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Stan...
func (s *server) startListentStanServer() {
	time.Sleep(time.Millisecond * 10)
	_, err := s.sc.Subscribe("orderDemonstrationServer", func(m *stan.Msg) {
		// o, isValid := s.ReadDataWithValidation(string(m.Data))
		// if isValid {
		// 	fmt.Printf("[%s]Recieved a message: %s\n", string(time.Now().Format("2006-01-02 15:04:05")), o.Id)
		// 	s.cache.Set(o.Id, o.Info)
		// }
		o := s.ReadData(string(m.Data))
		fmt.Printf("[%s]Recieved a message: %s\n", string(time.Now().Format("2006-01-02 15:04:05")), o.Id)
		s.cache.Set(o.Id, o.Info)

	}, stan.DeliverAllAvailable(), stan.DurableName(""), stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
		return
	}
	select {}
}

// Parsing string data into model.DatabaseData
func (s *server) ReadDataWithValidation(data string) (*model.DatabaseData, bool) {
	fmt.Println(data)
	jsonData := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	uid := jsonData["Id"].(string)
	delete(jsonData, "Id")
	o := model.DatabaseData{
		Id:   uid,
		Info: jsonData,
	}
	s.store.Insert(&o)
	return &o, o.Validate()
}

func (s *server) ReadData(data string) *model.DatabaseData {
	fmt.Println(data)
	jsonData := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	uid := jsonData["Id"].(string)
	delete(jsonData, "Id")
	o := model.DatabaseData{
		Id:   uid,
		Info: jsonData,
	}
	s.store.Insert(&o)
	return &o
}
