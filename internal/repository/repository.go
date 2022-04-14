package repository

import "github.com/amirhossein-ka/DigiFossGo/internal/model"

type Database interface {
	CreateSetting(chatID int64, setting *model.Setting) error
	ReadSetting(chatID int64) (*model.Setting, error)
	UpdateSetting(chatID int64, setting *model.Setting) error
	DefaultSetting(chatID int64) error
}
