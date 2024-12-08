package config

import (
	"log"
	"my-template/storage"
)

type Config struct {
	ConfigDB *storage.Config
	Storage  *storage.Storage
}

func NewConfig() *Config {
	return &Config{
		ConfigDB: storage.NewStorageConfig(),
	}
}

func (c *Config) ConfigStorageField() {
	storageConf := storage.NewStorage(c.ConfigDB)

	if err := storageConf.Open(); err != nil {
		log.Fatalf("not open db: %s", err)
	}

	c.Storage = storageConf
}
