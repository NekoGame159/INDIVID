package main

import ("database/sql"
		"fmt"
		"io/ioutil"
		"os"

		_"github.com/mattn/go-sqlite3" )

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createDBFile(dbName string) {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		_, err := os.Create(dbName)
		check(err)
	}
}

func main() {
	dbName := "test.db"
	sqlFile := "create_table.sql"
	createDBFile(dbName)

	db, err := sql.Open("sqlite3", dbName)
	check(err)
	defer db.Close()

	sqlFileContent, err := ioutil.ReadFile(sqlFile)
	check(err)

	_, err = db.Exec(string(sqlFileContent))
	check(err)
	
	fmt.Println("table created")
}