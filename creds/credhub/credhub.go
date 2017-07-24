package credhub

import (
	"path"

	"code.cloudfoundry.org/lager"

	"github.com/cloudfoundry-incubator/credhub-cli/api"
	"github.com/cloudfoundry-incubator/credhub-cli/models"
	"github.com/cloudfoundry/bosh-cli/director/template"
)

type Credhub struct {
	CredhubApi *api.Api

	logger lager.Logger

	PathPrefix   string
	TeamName     string
	PipelineName string
}

func (v Credhub) Get(varDef template.VariableDefinition) (interface{}, bool, error) {
	var secret models.CredentialResponse
	var found bool
	var err error

	if v.PipelineName != "" {
		secret, found, err = v.findSecret(v.path(v.TeamName, v.PipelineName, varDef.Name))
		if err != nil {
			return nil, false, err
		}
	}

	if !found {
		secret, found, err = v.findSecret(v.path(v.TeamName, varDef.Name))
		if err != nil {
			return nil, false, err
		}
	}

	if !found {
		return nil, false, nil
	}

	val, found := secret.ResponseBody["value"]
	if !found {
		return val, false, nil
	}

	return val, true, nil
}

func (v Credhub) findSecret(path string) (models.CredentialResponse, bool, error) {
	var secret models.CredentialResponse
	var err error

	_, err = v.CredhubApi.FindByPath(path)
	if err != nil {
		return secret, false, nil
	}

	secret, err = v.CredhubApi.GetByPath(path)
	if err != nil {
		return secret, false, err
	}

	return secret, true, nil
}

func (v Credhub) path(segments ...string) string {
	return path.Join(append([]string{v.PathPrefix}, segments...)...)
}

func (v Credhub) List() ([]template.VariableDefinition, error) {
	// Don't think this works with vault.. if we need it to we'll figure it out
	// var defs []template.VariableDefinition

	// secret, err := v.vaultClient.List(v.PathPrefix)
	// if err != nil {
	// 	return defs, err
	// }

	// var def template.VariableDefinition
	// for name, _ := range secret.Data {
	// 	defs := append(defs, template.VariableDefinition{
	// 		Name: name,
	// 	})
	// }

	return []template.VariableDefinition{}, nil
}
