// Package beacon-chain defines the entire runtime of an Ethereum beacon node.
package main

import (
	"context"
	"os"
	runtimeDebug "runtime/debug"

	"github.com/OffchainLabs/prysm/v6/cmd/beacon-chain/app"
	_ "github.com/OffchainLabs/prysm/v6/runtime/maxprocs"
	"github.com/sirupsen/logrus"
)

func main() {
	// rctx = root context with cancellation.
	// note other instances of ctx in this func are *cli.Context.
	rctx, cancel := context.WithCancel(context.Background())
	app := app.NewApp(cancel)

	log := logrus.WithField("prefix", "main")
	defer func() {
		// TODO can probably move to app.After, since it's called on panic also
		if x := recover(); x != nil {
			log.Errorf("Runtime panic: %v\n%v", x, string(runtimeDebug.Stack()))
			panic(x) // lint:nopanic -- This is just resurfacing the original panic.
		}
	}()

	if err := app.RunContext(rctx, os.Args); err != nil {
		log.Error(err.Error())
	}
}
