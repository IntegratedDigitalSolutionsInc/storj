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
	clusterPath          = "/Users/bohdanbashynskyi/storj-cluster"
	defaultDbEndpoint    = "postgresql://root@localhost:26257/metainfo?sslmode=disable"
	defaultSharedValues  = 0.3
	defaultBatchSize     = 10
	defaultWorkersNumber = 1
	defaultTotlaRecords  = 10
	defaultMetasearchAPI = "http://localhost:9998/meta_search"
)

// main parameters decalaration
var (
	dbEndpoint    string
	sharedValues  float64 = 0.3
	batchSize     int
	workersNumber int
	totalRecords  int
	mode          string
)

func readArgs() {
	flag.StringVar(&dbEndpoint, "db", defaultDbEndpoint, fmt.Sprintf("db endpoint, default: %v", defaultDbEndpoint))
	flag.StringVar(&mode, "md", metagenerator.DryRunMode, fmt.Sprintf("incert mode [%s, %s, %s], default: %v", metagenerator.ApiMode, metagenerator.DbMode, metagenerator.DryRunMode, metagenerator.DryRunMode))
	flag.Float64Var(&sharedValues, "sv", defaultSharedValues, fmt.Sprintf("percentage of shared values, default: %v", defaultSharedValues))
	flag.IntVar(&batchSize, "bs", defaultBatchSize, fmt.Sprintf("number of records per batch, default: %v", defaultBatchSize))
	flag.IntVar(&workersNumber, "wn", defaultWorkersNumber, fmt.Sprintf("number of workers, default: %v", defaultWorkersNumber))
	flag.IntVar(&totalRecords, "tr", defaultTotlaRecords, fmt.Sprintf("total number of records, default: %v", defaultTotlaRecords))
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
		sharedValues,  // 30% shared values
		batchSize,     // batch size
		workersNumber, // number of workers
		totalRecords,
		metagenerator.GetPathCount(ctx, db), // get path count
		metagenerator.GetProjectId(ctx, db).String(),
		os.Getenv("API_KEY"),
		mode, // incert mode
		defaultMetasearchAPI,
	)

	// Generate and insert/debug records
	startTime := time.Now()

	if err := batchGen.GenerateAndInsert(totalRecords); err != nil {
		panic(fmt.Sprintf("failed to generate records: %v", err))
	}

	fmt.Printf("Generated %v records in %v\n", totalRecords, time.Since(startTime))
}
