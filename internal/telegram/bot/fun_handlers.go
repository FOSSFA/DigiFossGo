package bot

import (
	"fmt"
	api "github.com/amirhossein-ka/DigiFossGo/pkg/adhan"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator/messages"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
)

func (h *handler) adhan(c tele.Context) error {
	var rem time.Duration
	msg := strings.Split(c.Message().Text, " ")[1:]
	m := strings.Join(msg, " ")
	if m == "" {
		m = "تهران"
	}
	rem, err := api.Reminder(m, strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		_ = c.Reply(h.translator.Translate(messages.GotAnError, translator.FA))
		return err
	}

	if rem.Seconds() < 0 {
		return c.Reply(h.translator.Translate(messages.AzanSohde, translator.FA))
	}

	ms := h.translator.Translate(messages.UntilAzan, translator.FA)
	dur := fmtDuration(rem)
	if err = c.Reply(fmt.Sprintf(ms, dur)); err != nil {
		return err
	}
	return nil
}
