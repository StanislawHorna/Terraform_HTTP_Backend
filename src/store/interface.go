package store

import (
	"terraform_http_backend/src/model"
)

type Store interface {
	GetState(projectName string, environment string) *model.TFState
	SaveState(projectName string, environment string, state model.TFState) error
}
