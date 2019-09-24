package event

import (
	"github.com/Daniil129932/goVkBot/vk/object"
)

type Command struct {
	Command        string
	Args           []string
	PrivateMessage *object.PrivateMessage
}

func (c *Command) GetName() string {
	return CommandEvent
}
