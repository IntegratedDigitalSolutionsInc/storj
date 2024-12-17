package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
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
	defaultApiKey        = "122VGqUbmpzC3cjoMkApS8vYjNNzqUGE6aHm6mwhhKBswTQugNxuxBF9zfof2h3h4Hw41u3hew3i1UpSGks7FHaxwkhW4ZSYRKQG61RK6haWgq3tAzDMQHqVEuyy9yXLo1RJnLzn"
	defaultClean         = false
)

// main parameters decalaration
var (
	dbEndpoint    string
	batchSize     int
	workersNumber int
	totalRecords  int
	mode          string
	apiKey        string
	clean         bool
)

func readArgs() {
	flag.StringVar(&dbEndpoint, "db", defaultDbEndpoint, fmt.Sprintf("db endpoint, default: %v", defaultDbEndpoint))
	flag.StringVar(&mode, "mode", metagenerator.DryRunMode, fmt.Sprintf("incert mode [%s, %s, %s], default: %v", metagenerator.ApiMode, metagenerator.DbMode, metagenerator.DryRunMode, metagenerator.DryRunMode))
	flag.StringVar(&apiKey, "apiKey", defaultApiKey, fmt.Sprintf("satelite api key, default: %v", defaultApiKey))
	flag.IntVar(&batchSize, "batchSize", defaultBatchSize, fmt.Sprintf("number of records per batch, default: %v", defaultBatchSize))
	flag.IntVar(&workersNumber, "workersNumber", defaultWorkersNumber, fmt.Sprintf("number of workers, default: %v", defaultWorkersNumber))
	flag.IntVar(&totalRecords, "totalRecords", defaultTotlaRecords, fmt.Sprintf("total number of records, default: %v", defaultTotlaRecords))
	flag.BoolVar(&clean, "clean", defaultClean, fmt.Sprintf("clean db before incerting, default: %v", defaultClean))
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

	if clean {
		metagenerator.CleanTable(ctx, db)
	}

	pathCount := metagenerator.GetPathCount(ctx, db)
	fmt.Printf("Detected %v records\n", pathCount)

	// Initialize batch generator
	batchGen := metagenerator.NewBatchGenerator(
		db,
		batchSize,     // batch size
		workersNumber, // number of workers
		totalRecords,
		pathCount, // get path count
		apiKey,
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
