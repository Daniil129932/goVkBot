package event

import (
	"github.com/Daniil129932/goVkBot/vk/object"
)

type MessageEdit struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}