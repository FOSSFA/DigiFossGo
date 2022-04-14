package bot

import tele "gopkg.in/telebot.v3"

func (b *bot) handlers() {
	f := b.bot.Group() // farsi handlers
	f.Use(b.handler.farsi)

	f.Handle(tele.OnText, nil)

	b.bot.Handle("/start", b.handler.start)
	b.bot.Handle("/hungry", b.handler.adhan)
	b.bot.Handle("/h", b.handler.adhan)
	b.bot.Handle("/help", b.handler.help)
}
