package bot

import (
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator/messages"
	tele "gopkg.in/telebot.v3"
)

func (h *handler) start(c tele.Context) error {
	// TODO change translator.FA from what seted in db
	if err := c.Send(h.translator.Translate(messages.Start, translator.FA)); err != nil {
		return err
	}
	return nil
}

func (h *handler) help(c tele.Context) error {
	if err := c.Send(h.translator.Translate(messages.Help, translator.FA)); err != nil {
		return err
	}

	return nil
}
