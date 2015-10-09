package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/willowtreeapps/timestamp"
	"github.com/willowtreeapps/timestamp/db"
)

var (
	dbPath string
)

func init() {
	flag.StringVar(&dbPath, "db-path", "testdb.sqlite",
		"Path at which to store test database")
}

func main() {
	flag.Parse()

	os.Remove(dbPath)
	db, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	if err := db.SetUp(); err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	if err := db.SeedData(); err != nil {
		log.Fatalf("Could not seed database: %v", err)
	}

	var timestamps []timestamp.Timestamp
	if err := db.AllTimestamps(&timestamps); err != nil {
		log.Fatalf("Could not fetch timestamps: %v", err)
	}
	fmt.Println("Found timestamps:")
	for _, t := range timestamps {
		fmt.Printf("  %v\n", time.Unix(t.SecondsSinceEpoch, 0).Format(time.UnixDate))
	}
}
