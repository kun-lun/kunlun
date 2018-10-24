package commands

import (
	"github.com/kun-lun/kunlun/common/fileio"
	"github.com/kun-lun/kunlun/common/storage"
	patching "github.com/kun-lun/kunlun/patching"
	infra "github.com/kun-lun/kunlun/producers/infra-producer"
	"github.com/kun-lun/kunlun/producers/infra-producer/handler"
)

type PlanInfra struct {
	stateStore storage.Store
	fs         fileio.Fs
	logger     logger
}

func NewPlanInfra(
	stateStore storage.Store,
	fs fileio.Fs,
	logger logger,
) PlanInfra {
	return PlanInfra{
		stateStore: stateStore,
		fs:         fs,
		logger:     logger,
	}
}

func (p PlanInfra) CheckFastFails(args []string, state storage.State) error {
	return nil
}

func (p PlanInfra) Execute(args []string, state storage.State) error {
	handlerType := handler.TerraformHandlerType // should get from args
	debug := true
	infraProducer, _ := infra.NewInfraProducer(p.stateStore, handlerType, debug)

	// load the provisioned manifest
	patching := patching.NewPatching(p.stateStore, p.fs)
	manifest, err := patching.ProvisionManifest()
	if err != nil {
		return err
	}

	return infraProducer.Setup(*manifest, state)
}
