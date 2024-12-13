package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"storj.io/storj/metagenerator"
)

// default values
const (
	defaultDbEndpoint    = "postgresql://root@localhost:26257/master?sslmode=disable"
	defaultSharedFields  = 0.3
	defaultBatchSize     = 10
	defaultWorkersNumber = 1
	defaultTotlaRecords  = 10
	defaultMetasearchAPI = "http://localhost:6666"
)

// main parameters decalaration
var (
	dbEndpoint    string
	batchSize     int
	workersNumber int
	totalRecords  int
	mode          string
)

func readArgs() {
	flag.StringVar(&dbEndpoint, "db", defaultDbEndpoint, fmt.Sprintf("db endpoint, default: %v", defaultDbEndpoint))
	flag.StringVar(&mode, "mode", metagenerator.DryRunMode, fmt.Sprintf("incert mode [%s, %s, %s], default: %v", metagenerator.ApiMode, metagenerator.DbMode, metagenerator.DryRunMode, metagenerator.DryRunMode))
	flag.IntVar(&batchSize, "batchSize", defaultBatchSize, fmt.Sprintf("number of records per batch, default: %v", defaultBatchSize))
	flag.IntVar(&workersNumber, "workersNumber", defaultWorkersNumber, fmt.Sprintf("number of workers, default: %v", defaultWorkersNumber))
	flag.IntVar(&totalRecords, "totalRecords", defaultTotlaRecords, fmt.Sprintf("total number of records, default: %v", defaultTotlaRecords))
	flag.Parse()
}

func main() {
	readArgs()

	// Connect to CockroachDB
	db, err := sql.Open("postgres", dbEndpoint)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	defer db.Close()
	ctx := context.Background()

	// Initialize batch generator
	batchGen := metagenerator.NewBatchGenerator(
		db,
		batchSize,     // batch size
		workersNumber, // number of workers
		totalRecords,
		metagenerator.GetPathCount(ctx, db), // get path count
		os.Getenv("API_KEY"),
		mode, // incert mode
		defaultMetasearchAPI,
		dbEndpoint,
	)

	// Generate and insert/debug records
	startTime := time.Now()

	if err := batchGen.GenerateAndInsert(ctx); err != nil {
		panic(fmt.Sprintf("failed to generate records: %v", err))
	}

	fmt.Printf("Generated %v records in %v\n", totalRecords, time.Since(startTime))
}
