package main

import (
	desmosapp "github.com/desmos-labs/desmos/app"
	desmosdb "github.com/desmos-labs/djuno/database"
	"github.com/desmos-labs/djuno/x"
	junocmd "github.com/desmos-labs/juno/cmd"
	"github.com/desmos-labs/juno/cmd/parse"
)

func main() {
	// Configure the parse command
	config := parse.NewConfig("djuno").
		WithRegistrar(x.NewModulesRegistrar()).
		WithEncodingConfigBuilder(desmosapp.MakeTestEncodingConfig).
		WithDBBuilder(desmosdb.Builder)

	// Run the commands and panic on any error
	executor := junocmd.BuildDefaultExecutor(config)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
