package model

import (
	"fmt"
	"time"
)

type (
	Group struct {
		GroupID uint    `json:"groupID,omitempty"`
		Name    string  `json:"name,omitempty"`
		ChatID  int64   `json:"chatID,omitempty"`
		Setting Setting `json:"setting"`
		Users   []Users `json:"users"`
	}

	Setting struct {
		MaxWarns      uint          `json:"maxWarns,omitempty"`
		MuteAfterWarn bool          `json:"muteAfterWarn"`
		TimeToMute    time.Duration `json:"timeToMute,omitempty"`
		Lang          string        `json:"lang"`
	}

	Users struct {
		UserID       int64        `json:"userID,omitempty"`
		Name         string       `json:"name"`
		Username     string       `json:"username"`
		Points       uint         `json:"points,omitempty"`
		Restrictions Restrictions `json:"restrictions"`
	}

	Restrictions struct {
		Muted      bool  `json:"muted,omitempty"`
		MutedUntil int64 `json:"mutedUntil"` // unix time
		Warns      uint8 `json:"warns,omitempty"`
		Banned     bool  `json:"banned"`
	}
)

func x() {
	fmt.Println(time.Now().Unix())
}
