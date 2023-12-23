package main

import (
	"os"
 	"log"
	 "database/sql"
	 "io/ioutil"

	 _"github.com/mattn/go-sqlite3"
)


func main() {
	filePath := "database.db"


	os.Remove(filePath)

	log.Printf("Creating %s", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Printf("%s created", filePath)

	sqliteDataBase, _ = sql.Open("sqlite3", filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	fileSql, err := ioutil.ReadFile("./bd.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(fileSql)

	sqliteDataBase.Close()

	if _, err := sqliteDataBase.Exec(string(fileSql));  err != nil {
		log.Fatal(err.Error())
	}

	// _, err = gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

}