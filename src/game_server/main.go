package main

import (
	"flag"
	_ "game_server/agent_store"
	"game_server/server"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Debug flags
var DEBUG_DBLOG = false
var DEBUG_HTTPLOG = false
var DEBUG_DBTODISK = false

func parseCommandline() {
	flag.BoolVar(&DEBUG_DBLOG, "dblog", false, "Log database transactions to stdout")
	flag.BoolVar(&DEBUG_DBTODISK, "ondiskdb", false, "Write sqlite database to disk")
	flag.BoolVar(&DEBUG_HTTPLOG, "httplog", false, "Log HTTP requests to stdout")

	flag.Parse()
}

func blockUntilInterrupt() {
	// Accept SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until shutdown signal recieved
	<-c
}

func main() {
	log.Println("Starting controller on '0.0.0.0:8080'...")

	parseCommandline()

	// Starts serving http in seperate goroutine
	gameServer := server.StartHTTPServer(":8080", DEBUG_HTTPLOG)

	//  --  Startup section over!  --  //
	log.Println("Controller ready!")

	// Block main thread until interrupt from OS
	blockUntilInterrupt()

	// Start shutdown once interrupt triggered
	log.Println("Stopping Controller...")
	server.StopHTTPServer(gameServer, time.Second*15) // Wait for existing connections to finish, then stop http server

	log.Println("Controller stopped, goodbye!")
	os.Exit(0)
}
