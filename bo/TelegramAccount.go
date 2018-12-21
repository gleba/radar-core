package bo

import "github.com/gleba/radar-core/ns"

func init() {
	ns.TelegramAccount.Register(TelegramAccount{})
}

type TelegramAccount struct {
	ID        int      `reindex:"id,,pk"` // 'id' is primary key
	FirstName string   // add index by 'name' field
	LastName  string   // add index by 'name' field
	UserName  string   `reindex:"user_name"` // add index by 'name' field
	Photos    []string // add index by articles 'articles' array
	//Runes []  *DeviceRune `reindex:"devices,,joined"` // add sortable index by 'year' field
}
