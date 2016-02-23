package main

import (
	"errors"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLないぞ？")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
