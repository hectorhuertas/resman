package main

import "fmt"

import "log"
import "os"
import  "resource-manager/gitLocal"
import "database/sql"
import _ "github.com/lib/pq"

func main() {
	// scan()
	db, err := sql.Open("postgres",
		"user=postgres dbname=pqgotest sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		// do something here
	}
	defer db.Close()
}

func scan()  {
	CODEPATH := os.Getenv("HOME") + "/xdev" // Might file for / dir
	fmt.Printf("CODEPATH: %s\n", CODEPATH)

	// fmt.Printf("All: %v\n", gitLocal.GetRepos(CODEPATH))
	// fmt.Printf("Clean: %v\n", gitLocal.GetUnmodifiedRepos(CODEPATH))
	fmt.Printf("Dirty: %v\n", gitLocal.GetModifiedRepos(CODEPATH))
}
