package model

type Order struct {
	order_uid          string    `json:"order_uid"`
	track_number       string    `json:"track_number"`
	entry              string    `json:"entry"`
	delivery           *Delivery `json:"delivery"`
	payment            *Payment  `json:"payment"`
	items              []Items   `json:"items"`
	locale             string    `json:"locale"`
	internal_signature string    `json:"internal_signature"`
	customer_id        string    `json:"customer_id"`
	delivery_service   string    `json:"delivery_service"`
	shardkey           string    `json:"shardkey"`
	sm_id              int       `json:"sm_id"`
	date_created       string    `json:"date_created"`
	oof_shard          string    `json:"oof_shard"`
}
