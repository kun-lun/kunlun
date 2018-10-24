package handler

import (
	artifacts "github.com/kun-lun/kunlun/artifacts"
	"github.com/kun-lun/kunlun/common/storage"
)

type Manager interface {
	Setup(manifest artifacts.Manifest, kunlunState storage.State) error
	Apply(kunlunState storage.State) (storage.State, error)
	GetOutputs() (Outputs, error)
}
