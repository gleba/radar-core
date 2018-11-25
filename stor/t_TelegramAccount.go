package stor

import (
	"github.com/restream/reindexer"
	"log"
)

const NSTelegramAccount string = "telegram_account"

type TelegramAccount struct {
	ID        int      `reindex:"id,,pk"` // 'id' is primary key
	FirstName string   // add index by 'name' field
	LastName  string   // add index by 'name' field
	UserName  string   `reindex:"user_name"` // add index by 'name' field
	Photos    []string // add index by articles 'articles' array
	//Runes []  *DeviceRune `reindex:"devices,,joined"` // add sortable index by 'year' field
}

func TelegramAccountFromID(id int) (acc *TelegramAccount, found bool) {
	item, found := DB.
		Query(NSTelegramAccount).
		Where("id", reindexer.EQ, id).
		Get()
	if found {
		acc = item.(*TelegramAccount)
		return acc, found
	} else {
		return nil, found
	}

}

func (acc TelegramAccount) Save() {
	log.Println("Save")
	err := DB.Upsert(NSTelegramAccount, acc)
	if err != nil {
		log.Fatal(acc)
		log.Fatal("↑", err)
	}
}
