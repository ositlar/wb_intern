package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
	stan "github.com/nats-io/stan.go"
	"github.com/ositlar/wb_intern_1/internal/apiserver"
	"github.com/ositlar/wb_intern_1/internal/model"
)

var (
	configPath string
)

const (
	clusterID  = "wbCluster"
	clientID   = "intern-sub"
	NatsSrvURL = "nats://127.0.0.1:4222"
)

func init() {
	flag.StringVar(&configPath, "config-flag", "configs/config.toml", "path yo config file")
}

func main() {
	flag.Parse()

	go testPublish()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

// Функция для теста работы Nats-streaming, отправляет примеры json-ов в канал
func testPublish() {
	sc, stanerr := stan.Connect(clusterID, clientID, stan.NatsURL(NatsSrvURL))
	if stanerr != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", stanerr, NatsSrvURL)
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(charset))
		result += string(charset[index])
	}
	testInfo := make(map[string]interface{})
	testInfo["test"] = "test"
	testOrder := model.DatabaseData{
		Id:   result,
		Info: testInfo,
	}
	test, err := json.Marshal(testOrder)
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("orderDemonstrationServer", test)
	time.Sleep(time.Millisecond * 1000)
	sc.Close()
}
