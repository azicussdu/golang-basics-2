package main

type Database interface {
	Connect() bool
	Query(query string) ([]map[string]interface{}, bool)
	Close() bool
}

type PostgresDB struct {
	isConnect bool
}

func (pdb *PostgresDB) Connect() bool {
	pdb.isConnect = true
	return true
}

func (pdb *PostgresDB) Close() bool {
	pdb.isConnect = false
	return true
}

func (pdb *PostgresDB) Query(query string) ([]map[string]interface{}, bool) {
	if pdb.isConnect {
		if len(query) > 0 {
			return []map[string]interface{}{
				{"id": 1, "name": "Arman", "height": 180.5},
				{"id": 2, "name": "Beka", "height": 185.5},
			}, true
		}
	}

	return nil, false
}

//func main() {
//	var db Database = &PostgresDB{}
//	db.Connect()
//	results, ok := db.Query("SELECT * FROM users")
//	if ok == false {
//		fmt.Println("Error: first connect to make queries")
//		return
//	}
//	fmt.Println("Query results:", results) //например: [ map[id:1 name:Alice] map[id:2 name:Bob] ]
//	db.Close()
//}
