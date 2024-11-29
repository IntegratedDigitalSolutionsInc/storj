package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"storj.io/storj/metagenerator"
	"storj.io/storj/metasearch"
)

const (
	apiKey    = "15XZjcVqxQeggDyDpPhqJvMUB6NtQ1CiuW6mAwzRAVNE5gtr7Yh12MdtqvVbYQ9rvCadeve1f2LGiB53QnFyVV9CTY5HAv3jtFvtnKiVvehh4Dz9jwYx6yhV5bD1wGBrADuKCkQxa"
	projectId = "9088e8cc-d344-4767-8e07-901abc2734b6"
)

var tRs = []int{
	100_000,
	900_000,
	9_000_000,
	99_000_000,
}

func setupSuite() (func(tb testing.TB), *sql.DB, context.Context) {
	// Connect to CockroachDB
	db, err := sql.Open("postgres", defaultDbEndpoint)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	ctx := context.Background()

	// Return a function to teardown the test
	return func(tb testing.TB) {
		metagenerator.CleanTable(ctx, db)
		db.Close()
	}, db, ctx
}

func BenchmarkSimpleQuery(b *testing.B) {
	teardownSuite, db, ctx := setupSuite()
	defer teardownSuite(b)
	for _, tR := range tRs {
		metagenerator.GeneratorSetup(1000, 10, tR, apiKey, projectId, defaultMetasearchAPI, db, ctx)
		for _, n := range metagenerator.MatchingEntries {
			if tR < n {
				break
			}
			b.Run(fmt.Sprintf("total_objects_%v_matching_entries_%d", tR, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					val := fmt.Sprintf("benchmarkValue_%v", n)
					b.ResetTimer()
					url := fmt.Sprintf("%s/metasearch/%s", defaultMetasearchAPI, metagenerator.Label)
					_, err := metagenerator.SearchMeta(metasearch.SearchRequest{
						Match: map[string]any{
							"field_" + val: val,
						},
					}, apiKey, projectId, url)

					if err != nil {
						panic(err)
					}
					/*
						b.StopTimer()
						var resp metagenerator.Response
						err = json.Unmarshal(res, &resp)
						if err != nil {
							panic(err)
						}
						fmt.Printf("Got %v entries\n", len(resp.Results))
					*/
				}
			})
		}
	}
}
