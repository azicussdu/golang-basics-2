package main

import (
	"errors"
)

type Database interface {
	Connect() error
	Query(query string) ([]map[string]interface{}, error)
	Close() error
}

type PostgresDB struct {
	isConnect bool
}

func (pdb *PostgresDB) Connect() error {
	pdb.isConnect = true
	return nil
}

func (pdb *PostgresDB) Close() error {
	pdb.isConnect = false
	return nil
}

func (pdb *PostgresDB) Query(query string) ([]map[string]interface{}, error) {
	if !pdb.isConnect {
		return nil, errors.New("it is not connected")
	}
	if len(query) > 0 {
		return []map[string]interface{}{
			{"id": 1, "name": "Arman", "height": 180.5},
			{"id": 2, "name": "Beka", "height": 185.5},
		}, nil
	} else {
		return nil, errors.New("query is empty")
	}
}

//func main() {
//	var db Database = &PostgresDB{}
//	err := db.Connect()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	results, err := db.Query("SELECT * FROM users")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("Query results:", results) //например: [ map[id:1 name:Alice] map[id:2 name:Bob] ]
//	err = db.Close()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
