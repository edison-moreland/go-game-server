package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"path/filepath"
)

const DBPATH = "game_db.sqlite3"

var databaseConnection *gorm.DB
var isDatabaseInitialized bool = false

func migrateModels(db *gorm.DB) {

	db.AutoMigrate(&Agent{})
}

// StartDB initialize Sqlite database
func StartDB(writeToDisk bool, logTransactions bool) *gorm.DB {
	if isDatabaseInitialized {
		log.Panic("Database can't be initialized twice!")
	}

	dbLocation := ":memory:"

	if writeToDisk {
		fullPath, err := filepath.Abs(DBPATH)
		if err != nil {
			log.Panic("Couldn't get absolute path for db")
		}

		// Try deleting old db, ignoring errors that probably mean old db doesn't exist
		os.Remove(fullPath)

		dbLocation = fullPath
		log.Printf("Database path: {%v}", fullPath)
	}

	db, err := gorm.Open("sqlite3", dbLocation)
	if err != nil {
		log.Panicf("Failed to connect database! {%#v}", err)
	}

	db.LogMode(logTransactions)
	migrateModels(db)

	databaseConnection = db
	isDatabaseInitialized = true

	return db
}

// StopDB closes connection to database
func StopDB() error {
	if isDatabaseInitialized != true {
		return errors.New("Database conection has not been initialized and can not be used")
	}

	return databaseConnection.Close()
}

// DBConnection returns the gorm connection manager
func DBConnection() (*gorm.DB, error) {
	if isDatabaseInitialized != true {
		return nil, errors.New("Database conection has not been initialized and can not be used")
	}
	return databaseConnection, nil
}
