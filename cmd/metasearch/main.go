// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"

	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/metasearch"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/satellitedb"
)

var runCfg metasearch.Config

func init() {
	runCfg.Read()
}

func main() {
	ctx := context.Background()
	log, err := runCfg.Log.Build()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	db, err := satellitedb.Open(ctx, log.Named("db"), runCfg.Database, satellitedb.Options{
		ApplicationName: "metadata-api",
	})
	if err != nil {
		log.Error("Error starting master database on metadata api:", zap.Error(err))
		return
	}
	defer func() {
		err = errs.Combine(err, db.Close())
	}()

	metabase, err := metabase.Open(ctx, log.Named("metabase"), runCfg.Metainfo.DatabaseURL,
		runCfg.Metainfo.Metabase("metasearch-api"),
	)
	if err != nil {
		log.Error("Error creating metabase connection on metadata api:", zap.Error(err))
		return
	}
	defer func() {
		err = errs.Combine(err, metabase.Close())
	}()

	repo := metasearch.NewMetabaseSearchRepository(metabase)
	auth := metasearch.NewHeaderAuth(db)
	metadataAPI, err := metasearch.NewServer(log, repo, auth, runCfg.Endpoint)
	if err != nil {
		log.Error("Error creating metadata api:", zap.Error(err))
		return
	}
	err = metadataAPI.Run()
	if err != nil {
		log.Error("Error running metadata api:", zap.Error(err))
		return
	}
}
