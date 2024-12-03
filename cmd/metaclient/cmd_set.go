// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/zeebo/clingy"
	"storj.io/storj/cmd/uplink/ulloc"
)

type cmdSet struct {
	access   *AccessOptions
	location string

	bucket string
	key    string

	inputfile string
	inputdata string

	metadata map[string]interface{}
}

func newCmdSet() *cmdSet {
	return &cmdSet{
		access: newAccessOptions(),
	}
}

func (c *cmdSet) Setup(params clingy.Parameters) {
	c.access.Setup(params)
	c.inputfile = params.Flag("input-file", "File containing metadata to set", "", clingy.Short('i')).(string)
	c.inputdata = params.Flag("data", "Metadata to set", "", clingy.Short('d')).(string)

	c.location = params.Arg("location", "Location of object (sj://BUCKET/KEY)").(string)
}

func (c *cmdSet) Validate() (err error) {
	err = c.access.Validate()
	if err != nil {
		return err
	}

	loc, err := ulloc.Parse(c.location)
	if err != nil {
		return fmt.Errorf("invalid location '%s': %w", c.location, err)
	}

	var ok bool
	c.bucket, c.key, ok = loc.RemoteParts()
	if !ok {
		return fmt.Errorf("invalid location '%s': must be remote", c.location)
	}

	if c.bucket == "" || c.key == "" {
		return fmt.Errorf("invalid location '%s': both bucket and key must be provided", c.location)
	}

	if c.inputfile == "" && c.inputdata == "" {
		return fmt.Errorf("either --input-file or --data must be provided")
	}

	return nil
}

func (c *cmdSet) Execute(ctx context.Context) (err error) {
	err = c.Validate()
	if err != nil {
		return err
	}

	err = c.setMetadata(ctx)
	if err != nil {
		return err
	}

	client := newMetaSearchClient(c.access)
	err = client.SetObjectMetadata(ctx, c.bucket, c.key, c.metadata)
	if err != nil {
		return fmt.Errorf("cannot set metadata: %w", err)
	}

	return nil
}

func (c *cmdSet) setMetadata(ctx context.Context) (err error) {
	if c.inputfile == "-" {
		err = c.setMetadataFromStdin(ctx)
	} else if c.inputfile != "" {
		err = c.setMetadataFromFile(ctx)
	}
	if err != nil {
		return fmt.Errorf("error reading metadata: %w", err)
	}

	err = json.Unmarshal([]byte(c.inputdata), &c.metadata)
	if err != nil {
		return fmt.Errorf("invalid metadata: %w", err)
	}
	return nil
}

func (c *cmdSet) setMetadataFromStdin(ctx context.Context) (err error) {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("cannot read from stdin: %w", err)
	}

	c.inputdata = string(buf)
	return nil
}

func (c *cmdSet) setMetadataFromFile(ctx context.Context) (err error) {
	file, err := os.Open(c.inputfile)
	if err != nil {
		return fmt.Errorf("cannot open input file: %w", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("cannot read from input file: %w", err)
	}

	c.inputdata = string(buf)
	return nil
}
