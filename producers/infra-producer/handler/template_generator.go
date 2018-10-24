package handler

import (
	artifacts "github.com/kun-lun/kunlun/artifacts"
	"github.com/kun-lun/kunlun/common/storage"
)

type TemplateGenerator interface {
	GenerateTemplate(artifacts.Manifest, storage.State) string
}
