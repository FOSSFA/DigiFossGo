package service

import (
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"github.com/amirhossein-ka/DigiFossGo/internal/repository"
	"github.com/amirhossein-ka/DigiFossGo/internal/service/setting"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
)

type Service interface {
	setting.Setting
}

type service struct {
	setting.Setting
}

func New(cfg *config.Bot, db repository.Database, log logger.Logger) Service {
	s := setting.New(cfg, db, log)
	return service{
		s,
	}
}
