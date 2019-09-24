package event

import (
	"github.com/Daniil129932/goVkBot/vk/object"
)

type MessageReply struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageReply) GetName() string {
	return MessageReplyEvent
}