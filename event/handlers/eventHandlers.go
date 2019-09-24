package handlers

import "github.com/Daniil129932/goVkBot/event"

type EventHandler func(event.Event) bool

// messages
type CommandHandler func(args []string, command *event.Command) bool
type MessageNewHandler func(messageNew *event.MessageNew) bool
type MessageEditHandler func(edit *event.MessageEdit) bool
type MessageReplyHandler func(reply *event.MessageReply) bool
type MessageAllowHandler func(allow *event.MessageAllow) bool
type MessageDenyHandler func(deny *event.MessageDeny) bool