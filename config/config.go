package config

import (
	"errors"
)

var (
	ErrFailedToBindUser = errors.New("User型への変換に失敗しました")
	ErrLoginFailed = errors.New("ログインに失敗しました")
)
