package vk

import (
	"github.com/Daniil129932/goVkBot/vk/object"
)

type Response struct {
	Response interface{}
	Error    *object.ResponseError
	Raw      map[string]interface{}
}
