package setting

import (
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"github.com/amirhossein-ka/DigiFossGo/internal/repository"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
)

type Setting interface {
	CreateSetting(chatID int64) error
}

type setting struct {
	cfg    *config.Bot
	db     repository.Database
	logger logger.Logger
}

func (s *setting) CreateSetting(chatID int64) error {

	return nil
}

func New(cfg *config.Bot, db repository.Database, log logger.Logger) Setting {
	return &setting{
		cfg:    cfg,
		db:     db,
		logger: log,
	}
}
