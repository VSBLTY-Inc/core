package main

import (
	"os"

	_ "github.com/VSBLTY-Inc/core/data/expression/script"
	"github.com/VSBLTY-Inc/core/engine"
	"github.com/VSBLTY-Inc/core/support/log"
)

var (
	cfgJson       string
	cfgEngine     string
	cfgCompressed bool
	flogoEngine   engine.Engine
)

func init() {
	log.SetLogLevel(log.RootLogger(), log.ErrorLevel)

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine 3: %s", err.Error())
		os.Exit(1)
	}

	flogoEngine, err = engine.New(cfg, engine.ConfigOption(cfgEngine, cfgCompressed), engine.DirectRunner)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine 4: %s", err.Error())
		os.Exit(1)
	}
}
