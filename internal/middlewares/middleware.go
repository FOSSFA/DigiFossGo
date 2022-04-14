package middlewares

import tele "gopkg.in/telebot.v3"

func AdminsOnly(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		a, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return err
		}
		for _, v := range a {
			if v.User.ID == c.Message().Sender.ID {
				return next(c)
			}
		}
		return nil
	}
}

func CreatorOnly(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		a, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return err
		}
		for _, m := range a {
			if m.Role == tele.Creator {
				return next(c)
			}
		}
		return nil
	}
}

func IsReply(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Message().IsReply() {
			return next(c)
		}
		return nil
	}
}
