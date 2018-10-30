package storage

import "github.com/xplaceholder/common/fileio"

const (
	STATE_SCHEMA = 14
	STATE_FILE   = "kid-state.json"
)

type Store struct {
	dir         string
	fs          fs
	stateSchema int
}

type fs interface {
	fileio.FileWriter
	fileio.Remover
	fileio.AllRemover
	fileio.Stater
	fileio.AllMkdirer
	fileio.DirReader
}
