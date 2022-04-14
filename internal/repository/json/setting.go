package json

import (
	"errors"
	"fmt"
	"github.com/amirhossein-ka/DigiFossGo/internal/model"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
	"golang.org/x/exp/slices"
)

var (
	ErrChatIDNotFound = errors.New("chatID not found in database")
)

func (j *json) CreateSetting(chatID int64, setting *model.Setting) error {
	var g []model.Group
	err := j.decoder.Decode(&g)
	if err != nil {
		return err
	}

	index := groupIndex(&g, chatID)
	if index == -1 {
		return ErrChatIDNotFound
	}
	j.logger.Info(logger.LogField{Section: "repository/json", Function: "CreateSetting", Params: setting,
		Message: fmt.Sprintf("creating setting for chatID: %d, setting: %v", chatID, setting)})
	g[index].Setting = *setting

	return j.updateDB(&g)
}

func (j *json) ReadSetting(chatID int64) (*model.Setting, error) {
	var g []model.Group
	if err := j.decoder.Decode(&g); err != nil {
		return nil, err
	}

	index := groupIndex(&g, chatID)

	return &g[index].Setting, nil
}

func (j *json) UpdateSetting(chatID int64, setting *model.Setting) error {
	var g []model.Group
	err := j.decoder.Decode(&g)
	if err != nil {
		return err
	}

	index := groupIndex(&g, chatID)
	if index == -1 {
		return ErrChatIDNotFound
	}
	j.logger.Info(logger.LogField{Section: "repository/json", Function: "UpdateSetting", Params: setting,
		Message: fmt.Sprintf("updating setting for chatID: %d, setting: %v", chatID, setting)})
	g[index].Setting = *setting

	return j.updateDB(&g)
}

func (j *json) DefaultSetting(chatID int64) error {
	var g []model.Group
	if err := j.decoder.Decode(&g); err != nil {
		return err
	}
	index := groupIndex(&g, chatID)

	g[index].Setting = defaultSetting()
	return nil
}

func (j *json) updateDB(g *[]model.Group) error {
	if err := j.file.Truncate(0); err != nil {
		return err
	}
	if _, err := j.file.Seek(0, 0); err != nil {
		return err
	}
	return j.encoder.Encode(g)
}

func defaultSetting() model.Setting {
	return model.Setting{
		MaxWarns:      3,
		MuteAfterWarn: false,
		TimeToMute:    0,
		Lang:          "en",
	}
}

func groupIndex(g *[]model.Group, chatID int64) int {
	return slices.IndexFunc(*g, func(e model.Group) bool {
		return e.ChatID == chatID
	})

}
