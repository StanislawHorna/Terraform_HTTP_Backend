package store

import (
	"terraform_http_backend/src/model"
)

type Store interface {
	GetState(projectName string) *model.TFState
	SaveState(projectName string, state model.TFState) error
}
