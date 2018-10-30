package config

import (
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
	"github.com/xplaceholder/infra-producer/fileio"
	"github.com/xplaceholder/xplaceholder/application"
)

type logger interface {
	Println(string)
}

type fs interface {
	fileio.Stater
	fileio.TempFiler
	fileio.FileReader
	fileio.FileWriter
}

func NewConfig(logger logger, fs fs) Config {
	return Config{
		stateBootstrap: bootstrap,
		logger:         logger,
		fs:             fs,
	}
}

type Config struct {
	stateBootstrap StateBootstrap
	logger         logger
	fs             fs
}

func ParseArgs(args []string) (GlobalFlags, []string, error) {
	var globals GlobalFlags
	parser := flags.NewParser(&globals, flags.IgnoreUnknown)
	remainingArgs, err := parser.ParseArgs(args[1:])
	if err != nil {
		return GlobalFlags{}, remainingArgs, err
	}

	if !filepath.IsAbs(globals.StateDir) {
		workingDir, err := os.Getwd()
		if err != nil {
			return GlobalFlags{}, remainingArgs, err
		}
		globals.StateDir = filepath.Join(workingDir, globals.StateDir)
	}

	return globals, remainingArgs, nil
}

func (c Config) Bootstrap(globalFlags GlobalFlags, remainingArgs []string, argsLen int) (application.Configuration, error) {
	return application.Configuration{
		Command: "help",
	}, nil
}
