package apis

import (
	"github.com/kun-lun/kunlun/artifacts/deployments"
	"github.com/kun-lun/kunlun/common/fileio"
	"github.com/kun-lun/kunlun/common/storage"
	"github.com/kun-lun/kunlun/producers/deployment-producer/ashandler/generator"
)

type logger interface {
	Step(string, ...interface{})
	Printf(string, ...interface{})
	Println(string)
	Prompt(string) bool
}

type ASHandler struct {
	asGenerator generator.ASGenerator
}

func NewASHandler(
	stateStore storage.Store,
	logger logger,
	fs fileio.Fs,
) ASHandler {
	return ASHandler{
		asGenerator: generator.NewASGenerator(stateStore, logger, fs),
	}
}
func (a ASHandler) Handle(hostGroups []deployments.HostGroup, deployments []deployments.Deployment) error {
	return a.asGenerator.Generate(hostGroups, deployments)
}
