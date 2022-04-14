package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	cfg                     *Config
	ErrUnKnownFileExtension = errors.New("unknown file extension")
)

func Parse(path string, cfg *Config) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse config: %w", err)
		}
	}()
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseConfig(path, cfg)
	default:
		return ErrUnKnownFileExtension
	}

}

func parseConfig(path string, cfg *Config) (err error) {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
} 

func ParseEnv(cfg *Config) {
  envconfig.MustProcess("bot", &cfg.Bot)
}


func SetConfig(c *Config) {
	cfg = c
}
