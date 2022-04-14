package json

import (
	js "encoding/json"
	"errors"
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"github.com/amirhossein-ka/DigiFossGo/internal/repository"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
	"os"
	"strings"
)

var ErrUnknownDBType = errors.New("database type is not json")

type json struct {
	decoder *js.Decoder
	encoder *js.Encoder
	file    *os.File
	logger  logger.Logger
}

func New(cfg config.Database, log logger.Logger) (repository.Database, error) {
	file, err := newReader(cfg)
	if err != nil {
		log.Error(logger.LogField{Section: "repository", Function: "newReader()", Params: cfg, Message: err.Error()})
		return nil, err
	}
	decoder := js.NewDecoder(file)
	encoder := js.NewEncoder(file)
	encoder.SetIndent("", strings.Repeat(" ", cfg.JsonIndent))
	return &json{
		decoder: decoder,
		encoder: encoder,
		file:    file,
		logger:  log,
	}, err
}

func newReader(cfg config.Database) (*os.File, error) {
	// check database type
	switch cfg.Type {
	case "json":
		file, err := openFile(cfg.Path)
		if err != nil {
			return nil, err
		}
		return file, nil
	default:
		return nil, ErrUnknownDBType
	}
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}
