package sqlstore_test

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/ositlar/wb_intern_1/internal/model"
	"github.com/ositlar/wb_intern_1/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

var (
	testData = []byte(`{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
			"name": "Test Testov",
			"phone": "+9720000000",
			"zip": "2639809",
			"city": "Kiryat Mozkin",
			"address": "Ploshad Mira 15",
			"region": "Kraiot",
			"email": "test@gmail.com"
		},
		"payment": {
			"transaction": "b563feb7b2b84b6test",
			"request_id": "",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 1817,
			"payment_dt": 1637907727,
			"bank": "alpha",
			"delivery_cost": 1500,
			"goods_total": 317,
			"custom_fee": 0
		},
		"items": [
			{
				"chrt_id": 9934930,
				"track_number": "WBILMTESTTRACK",
				"price": 453,
				"rid": "ab4219087a764ae0btest",
				"name": "Mascaras",
				"sale": 30,
				"size": "0",
				"total_price": 317,
				"nm_id": 2389212,
				"brand": "Vivienne Sabo",
				"status": 202
			}
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`)
)

func TestInsert(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(charset))
		result += string(charset[index])
	}
	allInfo := make(map[string]interface{})
	deserErr := json.Unmarshal(testData, &allInfo)
	assert.NoError(t, deserErr)
	uid := result
	delete(allInfo, "order_uid")
	dbData := &model.DatabaseData{
		Id:   uid,
		Info: allInfo,
	}
	db, err := sql.Open("postgres", "host=localhost dbname=wb_intern_1 user=dev1 password=1a2s3d4f5g sslmode=disable")
	assert.NoError(t, err)
	store := sqlstore.NewStore(db)
	insErr := store.Insert(dbData)
	assert.NoError(t, insErr)
	db.Close()
}

func TestFindById(t *testing.T) {
	id := "b563feb7b2b84b6test"
	db, err := sql.Open("postgres", "host=localhost dbname=wb_intern_1 user=dev1 password=1a2s3d4f5g sslmode=disable")
	assert.NoError(t, err)
	store := sqlstore.NewStore(db)
	o, err := store.FindById(id)
	assert.NoError(t, err)
	assert.NotNil(t, o)
	db.Close()
}

func TestSelectAll(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost dbname=wb_intern_1 user=dev1 password=1a2s3d4f5g sslmode=disable")
	assert.NoError(t, err)
	store := sqlstore.NewStore(db)
	arr, selectErr := store.SelectAll()
	assert.NoError(t, selectErr)
	assert.NotNil(t, arr)
	db.Close()
}
