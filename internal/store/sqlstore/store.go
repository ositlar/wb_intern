package sqlstore

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq" //...
	"github.com/ositlar/wb_intern_1/internal/model"
)

type Store struct {
	db *sql.DB
	//orderRepository *OrderRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) FindById(id string) (*model.DatabaseData, error) {
	dbData := &model.DatabaseData{}
	var stringData string
	if err := s.db.QueryRow("SELECT id, info FROM orders WHERE id = $1", id).Scan(&dbData.Id, &stringData); err != nil {
		if err == sql.ErrNoRows {
			s.db.Close()
			return nil, err
		}
		s.db.Close()
		return nil, err
	}
	err := json.Unmarshal([]byte(stringData), &dbData.Info)
	if err != nil {
		s.db.Close()
		return nil, err
	}
	s.db.Close()
	return dbData, nil
}

func (s *Store) Insert(o *model.DatabaseData) error {
	jsonData, err := json.Marshal(o.Info)
	if err != nil {
		//s.db.Close()
		return err
	}
	if insertErr := s.db.QueryRow("INSERT INTO orders (id, info) VALUES ($1, $2)", o.Id, jsonData); insertErr != nil {
		//s.db.Close()
		return insertErr.Err()
	}
	//s.db.Close()
	return nil
}

func (s *Store) SelectAll() ([]model.DatabaseData, error) {

	rows, err := s.db.Query("SELECT id, info FROM orders")
	if err != nil {
		s.db.Close()
		return nil, err
	}

	defer rows.Close()

	var data []model.DatabaseData

	for rows.Next() {
		var id string
		var info string

		err = rows.Scan(&id, &info)
		if err != nil {
			s.db.Close()
			return nil, err
		}

		infoMap := make(map[string]interface{})

		err = json.Unmarshal([]byte(info), &infoMap)
		if err != nil {
			s.db.Close()
			return nil, err
		}
		data = append(data, model.DatabaseData{
			Id:   id,
			Info: infoMap,
		})
	}
	if err = rows.Err(); err != nil {
		s.db.Close()
		return nil, err
	}
	s.db.Close()
	return data, nil
}
