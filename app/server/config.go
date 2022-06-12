package server

import (
	"encoding/json"
	"os"

	"github.com/rtnhnd255/payment_service_emulator/app/storage"
)

type Config struct {
	Addr          string `json:"addr"`
	StorageConfig *storage.Config
}

func ParseConfig(path string, store *storage.Config) (result *Config, err error) {
	var c Config
	cfgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(cfgFile)
	jsonParser.Decode(&c)
	c.StorageConfig = store
	return &c, nil
}
