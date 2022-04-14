package telegram

type TeleBot interface {
	Start() error
	Stop()
}
