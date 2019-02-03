package modules

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq" // Postgres
)

func CheckDB() {

	// ( host -> ist der name des DB-Service aus der docker-compose.yml)
	//dbinfo := fmt.Sprintf("host=database port=5432 user=%s pass=%s dbname=%s sslmode=disable",
	//	_user, _pass, _dbname)

	//dbinfo := fmt.Sprintf("host=database port=5432 user=postgres dbname=postgres sslmode=disable",
	//	_user, _dbname)

	/*	DB_HOST := "database"
		DB_PORT := "5432"
		DB_USER := "postgres"
		DB_PASSWORD := "pgpass"
		DB_NAME := "appdb"
	*/

	fmt.Println("INfo: TestInfoComment: ", AppConfig.Info)

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)

	fmt.Printf("Host=%s Port=%s User=%s DBPasswd=%s DBName=%s\n",
		AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)

	fmt.Println("INFO: Try DB Connect...")
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("ERROR: DB Connection: ", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from testtable;")
	if err != nil {
		fmt.Println("ERROR: DB Select error: ", err.Error())
	}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		//fmt.Println("%3 | %8 ", id, name)
		fmt.Println(id, name)
	}

}
