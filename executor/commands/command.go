package commands

import (
	"github.com/kun-lun/kunlun/common/storage"
)

type CommandSet map[string]Command

type Command interface {
	CheckFastFails(subcommandFlags []string, state storage.State) error
	Execute(subcommandFlags []string, state storage.State) error
	Usage() string
}
