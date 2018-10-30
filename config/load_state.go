package config

import (
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
	"github.com/xplaceholder/common/fileio"
	"github.com/xplaceholder/common/storage"
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

func NewConfig(bootstrap storage.StateBootstrap, logger logger, fs fs) Config {
	return Config{
		stateBootstrap: bootstrap,
		logger:         logger,
		fs:             fs,
	}
}

type Config struct {
	stateBootstrap storage.StateBootstrap
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

	var command string
	if len(remainingArgs) > 0 {
		command = remainingArgs[0]
	}

	if len(remainingArgs) == 0 {
		return application.Configuration{
			Command: "help",
		}, nil
	}

	state, err := c.stateBootstrap.GetState(globalFlags.StateDir)
	if err != nil {
		return application.Configuration{}, err
	}

	return application.Configuration{
		Global: application.GlobalConfiguration{
			Debug:    globalFlags.Debug,
			StateDir: globalFlags.StateDir,
			Name:     globalFlags.EnvID,
		},
		State:                state,
		Command:              command,
		SubcommandFlags:      remainingArgs[1:],
		ShowCommandHelp:      false,
		CommandModifiesState: modifiesState(command),
	}, nil
}

func modifiesState(command string) bool {
	_, ok := map[string]struct{}{
		"digest":     {}, // detect the project type and generate the draft manifests.
		"plan-lift":  {}, // parse the draft manifests and generate the infrastructure manifests. (now in terraform)
		"lift":       {}, // run the infra manifests. prepare the environment.
		"plan-shift": {}, // generate the deployment scripts, (now in ansible)
		"shift":      {}, // run the deployment scripts
		"destroy":    {}, // destroy the environment we just setup.
	}[command]
	return ok
}
