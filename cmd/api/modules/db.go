package modules

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq" // Postgres
)

func CheckDB(_user string, _pass string, _dbname string) {

	// ( host -> ist der name des DB-Service aus der docker-compose.yml)
	dbinfo := fmt.Sprintf("host=database port=5432 user=%s pass=%s dbname=%s sslmode=disable",
		_user, _pass, _dbname)

	fmt.Printf("user=%s pass=%s dbname=%s sslmode=disable\n",
		_user, _pass, _dbname)

	fmt.Println("INFO: Try DB Connect...")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("ERROR: DB Connection: ", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from testtable;")
	if err != nil {
		fmt.Println("ERROR: DB Connection: ", err.Error())
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fmt.Println("%3v ")
	}

}
