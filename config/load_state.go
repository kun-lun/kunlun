package config

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/xplaceholder/infra-producer/fileio"
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
		logger: logger,
		fs:     fs,
	}
}

type Config struct {
	logger logger
	fs     fs
}

func ParseArgs(args []string) (GlobalFlags, []string, error) {
	var globals GlobalFlags
	parser := flags.NewParser(&globals, flags.IgnoreUnknown)
	remainingArgs, err := parser.ParseArgs(args[1:])
	if err != nil {
		return GlobalFlags{}, remainingArgs, err
	}
	return globals, remainingArgs, nil
}
