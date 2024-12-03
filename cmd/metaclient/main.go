// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zeebo/clingy"
)

func main() {
	ctx := context.Background()

	ok, err := clingy.Environment{
		Name: "metaclient",
	}.Run(ctx, func(cmds clingy.Commands) {
		commands(cmds)
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	if !ok || err != nil {
		os.Exit(1)
	}
}

func commands(cmds clingy.Commands) {
	cmds.New("get", "Get metadata for an existing object", newCmdGet())
}
