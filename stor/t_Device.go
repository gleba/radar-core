package stor

import "time"

const NSDeviceRune string = "device_rune"

type DeviceRune struct {
	Rune        string `reindex:"id,,pk"`
	Agent       string
	AccountsIDs []int              `reindex:"accounts_ids"`
	Accounts    []*TelegramAccount `reindex:"accounts,,joined"`
	Sessions    []string
	LastConnect time.Time
}

func (dev DeviceRune) HasAccountID(id int) bool {
	for _, accId := range dev.AccountsIDs {
		if accId == id {
			return true
		}
	}
	return false
}
