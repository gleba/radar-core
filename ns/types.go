package ns

import "github.com/gleba/radar-core/stor"

const (
	ChRuneAuth string = "rune.auth"
)

type RuneAuth struct {
	Rune    string
	Account *stor.TelegramAccount
}
