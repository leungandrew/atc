package credhub

import (
	"code.cloudfoundry.org/lager"

	"github.com/cloudfoundry-incubator/credhub-cli/api"
	"github.com/concourse/atc/creds"
)

type credhubFactory struct {
	credhubApi *api.Api
	logger     lager.Logger
	prefix     string
}

func NewCredhubFactory(logger lager.Logger, credhubApi *api.Api, prefix string) *credhubFactory {
	factory := &credhubFactory{
		credhubApi: credhubApi,
		logger:     logger,
		prefix:     prefix,
	}

	logger.Info("init-credhub", lager.Data{"api": credhubApi})
	return factory
}

func (factory *credhubFactory) NewVariables(teamName string, pipelineName string) creds.Variables {
	return &Credhub{
		CredhubApi:   factory.credhubApi,
		PathPrefix:   factory.prefix,
		TeamName:     teamName,
		logger:       factory.logger.Session("credhub-vars"),
		PipelineName: pipelineName,
	}
}
