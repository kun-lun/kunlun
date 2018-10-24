package apis

import (
	"github.com/kun-lun/kunlun/artifacts"

	"github.com/kun-lun/kunlun/artifacts/deployments"
	"github.com/kun-lun/kunlun/common/fileio"
	"github.com/kun-lun/kunlun/common/storage"
	ashandler "github.com/kun-lun/kunlun/producers/deployment-producer/ashandler"
	"github.com/kun-lun/kunlun/producers/deployment-producer/dpbuilder"
)

type logger interface {
	Step(string, ...interface{})
	Printf(string, ...interface{})
	Println(string)
	Prompt(string) bool
}

type DeploymentProducer struct {
	stateStore storage.Store
	logger     logger
	fs         fileio.Fs
}

func NewDeploymentProducer(
	stateStore storage.Store,
	logger logger,
	fs fileio.Fs,
) DeploymentProducer {
	return DeploymentProducer{
		stateStore: stateStore,
		logger:     logger,
		fs:         fs,
	}
}

type deploymentItem struct {
	hostGroup  deployments.HostGroup
	deployment deployments.Deployment
}

func (dp DeploymentProducer) Produce(
	manifest apis.Manifest,
) error {
	// generate the deployments
	dpBuilder := dpbuilder.DeploymentBuilder{}
	hostGroups, deployments, err := dpBuilder.Produce(manifest)
	if err != nil {
		return err
	}

	// generate the ansible scripts based on the deployments.
	asHandler := ashandler.NewASHandler(dp.stateStore, dp.logger, dp.fs)
	err = asHandler.Handle(hostGroups, deployments)
	if err != nil {
		return err
	}
	return nil
}
