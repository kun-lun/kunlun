package application

import (
	"fmt"

	"github.com/xplaceholder/common/errors"

	"github.com/xplaceholder/executor/commands"
)

type App struct {
	commands      commands.CommandSet
	configuration Configuration
}

func New(commands commands.CommandSet, configuration Configuration) App {
	return App{
		commands:      commands,
		configuration: configuration,
	}
}
func (a App) Run() error {
	err := a.execute()
	return err
}

func (a App) getCommand(commandString string) (commands.Command, error) {
	command, ok := a.commands[commandString]
	if !ok {
		// a.usage.Print()
		return nil, fmt.Errorf("unknown command: %s", commandString)
	}
	return command, nil
}

func (a App) execute() error {
	return &errors.NotImplementedError{}
}
