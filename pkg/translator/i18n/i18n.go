package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type messageBundle struct {
	bundle *i18n.Bundle
}

func New(path string) (translator.Translator, error) {
	bundle := &messageBundle{
		bundle: i18n.NewBundle(language.English),
	}
	if err := bundle.loadBundle(path); err != nil {
		return nil, err
	}

	return bundle, nil
}

func (m *messageBundle) loadBundle(path string) error {
	m.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	_, err := m.bundle.LoadMessageFile(path + "messages.en.toml")
	if err != nil {
		return err
	}
	_, err = m.bundle.LoadMessageFile(path + "messages.fa.toml")
	if err != nil {
		return err
	}

	return nil
}

func (m *messageBundle) getLocalized(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(m.bundle, lang)
}

func (m *messageBundle) TranslateEN(key string) string {
	message, err := m.getLocalized("en").Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		return key
	}
	return message
}

func (m *messageBundle) Translate(key string, lang translator.Lang) string {
	message, err := m.getLocalized(string(lang)).Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		return key
	}
	return message
}
