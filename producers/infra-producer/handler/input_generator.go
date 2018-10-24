package handler

import (
	artifacts "github.com/kun-lun/kunlun/artifacts"
	"github.com/kun-lun/kunlun/common/storage"
)

type InputGenerator interface {
	GenerateInput(artifacts.Manifest, storage.State) (map[string]interface{}, error)
	Credentials(state storage.State) map[string]string
}
