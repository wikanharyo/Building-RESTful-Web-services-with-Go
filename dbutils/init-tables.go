package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println("Driver Error!")
	}
	// Create table train
	_, statementErr := statement.Exec()
	if statementErr != nil {
		log.Println("Table already exists!")
	}
	statement, _ = dbDriver.Prepare(station)
	statement.Exec()
	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()
	log.Println("All tables created successfully")
}
