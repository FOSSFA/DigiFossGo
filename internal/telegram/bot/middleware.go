package bot

import (
	tele "gopkg.in/telebot.v3"
	"regexp"
)

func (h *handler) farsi(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		hungry := regexp.MustCompile(`^گشنمه.*`)
		teshneh := regexp.MustCompile(`^تشنمه.*`)
		if hungry.MatchString(c.Message().Text) || teshneh.MatchString(c.Message().Text) {
			return h.adhan(c)
		}
		return nil
	}
}
