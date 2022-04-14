package application

import (
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"github.com/amirhossein-ka/DigiFossGo/internal/repository/json"
	"github.com/amirhossein-ka/DigiFossGo/internal/service"
	"github.com/amirhossein-ka/DigiFossGo/internal/telegram"
	"github.com/amirhossein-ka/DigiFossGo/internal/telegram/bot"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger/zap"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator/i18n"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {
	tr, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	log := zap.New(cfg.Log.Path)

	db, err := json.New(cfg.Database, log)
	if err != nil {
		return err
	}

	s := service.New(&cfg.Bot, db, log)

	b, err := bot.New(cfg.Bot, s, log, tr)
	if err != nil {
		return err
	}

	go graceFullStop(log, b)

	if err = b.Start(); err != nil {
		return err
	}

	return nil
}

func graceFullStop(l logger.Logger, b telegram.TeleBot) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	// blocks until recived a signal
	<-sig

	l.Info(logger.LogField{Section: "internal/application", Function: "graceFullStop()", Params: nil,
		Message: "Received interrupt, closing connection"})

	b.Stop()
}
