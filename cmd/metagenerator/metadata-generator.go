package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"storj.io/common/macaroon"
	"storj.io/storj/metagenerator"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/satellitedb"
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

	rawToken := os.Getenv("API_KEY")
	var projectId string
	if mode == metagenerator.DbMode {
		projectId = getProjectId(rawToken)
	}

	// Initialize batch generator
	batchGen := metagenerator.NewBatchGenerator(
		db,
		batchSize,     // batch size
		workersNumber, // number of workers
		totalRecords,
		metagenerator.GetPathCount(ctx, db), // get path count
		projectId,
		rawToken,
		mode, // incert mode
		defaultMetasearchAPI,
	)

	// Generate and insert/debug records
	startTime := time.Now()

	if err := batchGen.GenerateAndInsert(ctx); err != nil {
		panic(fmt.Sprintf("failed to generate records: %v", err))
	}

	fmt.Printf("Generated %v records in %v\n", totalRecords, time.Since(startTime))
}

func getProjectId(rawToken string) string {
	ctx := context.Background()
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	db, err := satellitedb.Open(context.Background(), log.Named("db"), dbEndpoint, satellitedb.Options{
		ApplicationName: "metadata-api",
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		err = errs.Combine(err, db.Close())
	}()

	// Parse API token
	apiKey, err := macaroon.ParseAPIKey(rawToken)
	if err != nil {
		panic(err)
	}

	// Get projectId
	var keyInfo *console.APIKeyInfo
	keyInfo, err = db.Console().APIKeys().GetByHead(ctx, apiKey.Head())
	if err != nil {
		panic(err)
	}
	return keyInfo.ProjectID.String()
}
