package main

import (
	"log"
	"os"

	clogger "github.com/kun-lun/common/logger"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/executor/commands"
	executor "github.com/kun-lun/executor/pkg/apis"
	"github.com/kun-lun/kunlun/config"
	"github.com/spf13/afero"
)

var Version = "dev"

func main() {
	log.SetFlags(0)

	logger := clogger.NewLogger(os.Stdout, os.Stdin)
	stderrLogger := clogger.NewLogger(os.Stderr, os.Stdin)
	stateBootstrap := storage.NewStateBootstrap(stderrLogger, Version)

	globals, remainingArgs, err := config.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if globals.NoConfirm {
		logger.NoConfirm()
	}
	// stateJSON, _ := json.Marshal(globals)
	// stderrLogger.Println(string(stateJSON))

	// File IO
	fs := afero.NewOsFs()
	afs := &afero.Afero{Fs: fs}

	// Configuration

	stateStore := storage.NewStore(globals.StateDir, afs)
	stateMerger := config.NewMerger(afs)
	newConfig := config.NewConfig(stateBootstrap, stateMerger, stderrLogger, afs)

	appConfig, err := newConfig.Bootstrap(globals, remainingArgs, len(os.Args))
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	// // Utilities
	// envIDGenerator := helpers.NewEnvIDGenerator(rand.Reader)
	usage := commands.NewUsage(logger)

	app := executor.NewExecutor(appConfig, usage, logger, stateStore)

	err = app.Run()
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
