package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	db "github.com/vivek-344/token-management-system/db/queries"
	token "github.com/vivek-344/token-management-system/internal"
	"github.com/vivek-344/token-management-system/util"
)

func main() {
	// Load Config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Establish database connection
	conn, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer conn.Close()

	queries := db.NewQueries(conn)

	// Initialize tokens
	err = token.InitializeTokens(queries)
	if err != nil {
		log.Fatal("Error initializing tokens: ", err)
	}

	// Number of operations to simulate
	var operations int
	fmt.Print("Enter simulation time (operations): ")
	_, err = fmt.Scan(&operations)

	if err != nil {
		log.Fatal("Error:", err, "- Please enter an integer.")
		return
	}

	// Simulate operations
	err = token.SimulateOperations(queries, operations)
	if err != nil {
		log.Fatal("Error during simulation: ", err)
	}

	// Display results
	err = token.DisplayResults(queries)
	if err != nil {
		log.Fatal("Error displaying results: ", err)
	}
}
