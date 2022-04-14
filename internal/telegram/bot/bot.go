package bot

import (
	"github.com/amirhossein-ka/DigiFossGo/internal/config"
	"github.com/amirhossein-ka/DigiFossGo/internal/service"
	"github.com/amirhossein-ka/DigiFossGo/internal/telegram"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"net/url"
	"time"
)

type bot struct {
	bot     *tele.Bot
	handler *handler
}

type handler struct {
	logger     logger.Logger
	translator translator.Translator
	service    service.Service
	// TODO add service when completed
}

func New(cfg config.Bot, service service.Service, logger logger.Logger, tr translator.Translator) (telegram.TeleBot, error) {
	s, err := botSetting(cfg)
	if err != nil {
		return nil, err
	}
	b, err := tele.NewBot(s)
	if err != nil {
		return nil, err
	}

	return &bot{
		bot: b,
		handler: &handler{
			logger:     logger,
			translator: tr,
			service:    service,
		},
	}, nil
}

func (b *bot) Start() error {
	b.handlers()
	b.bot.Start()
	return nil
}

func (b *bot) Stop() {
	b.handler.logger.Info(logger.LogField{Section: "telegram/bot", Function: "Stop() error", Params: nil,
		Message: "Stopping bot..."})

	b.bot.Stop()
}

func defaultSetting(cfg config.Bot) tele.Settings {
	return tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: time.Second * 10},
	}
}

func botSetting(cfg config.Bot) (tele.Settings, error) {
	if cfg.UseProxy {
		u, err := url.Parse(cfg.ProxyAddress)
		if err != nil {
			return tele.Settings{}, err
		}
		return tele.Settings{
			Token:  cfg.Token,
			Poller: &tele.LongPoller{Timeout: time.Second * 10},
			Client: &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(u)}},
		}, nil
	} else {
		return defaultSetting(cfg), nil
	}
}
