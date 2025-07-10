package file

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"terraform_http_backend/src"
	"terraform_http_backend/src/log"
	"terraform_http_backend/src/model"
	"terraform_http_backend/src/store"
)

var (
	instance *FileStore
	once     sync.Once
)

type FileStore struct {
	path      string
	extension string
}

func (fs *FileStore) GetState(projectName string) *model.TFState {
	filePath := filepath.Join(fs.path, projectName+fs.extension)
	state := readFile(filePath)
	if state == nil {
		state = &model.TFState{Version: 1}
	}

	return state
}
func (fs *FileStore) SaveState(projectName string, state model.TFState) error {
	filePath := filepath.Join(fs.path, projectName+fs.extension)
	err := writeFile(filePath, state)
	if err != nil {
		return log.Error(err, "Writing file failed")
	}
	return nil
}

func newInstance(path string, fileExtension string) *FileStore {
	instance := FileStore{path: path, extension: fileExtension}

	// Create directory if it doesn't exist
	err := os.MkdirAll("states", 0755)
	if err != nil {
		return nil
	}
	return &instance
}

func GetInstance() *FileStore {
	once.Do(func() {
		conf := src.GetConfig()
		instance = nil
		if conf.StoreType == store.StoreType_file {
			instance = newInstance(conf.FileStore.Path, conf.FileStore.FileExtension)
		}

	})
	return instance
}

func readFile(path string) *model.TFState {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error(err, "Failed to read a file")
		return nil
	}

	var state model.TFState
	err = json.Unmarshal(data, &state)
	if err != nil {
		log.Error(err, "Failed to parse file")
		return nil
	}
	return &state
}

func writeFile(path string, state model.TFState) error {
	data, err := json.MarshalIndent(state, "", "    ")
	if err != nil {
		return log.Error(err, "Failed to marshal JSON")
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return log.Error(err, "Failed to write a file")
	}
	return nil
}
