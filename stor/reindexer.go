package stor

import (
	"github.com/restream/reindexer"
	"log"
	"os"
)

var DB *reindexer.Reindexer

func Connect() {
	DB = reindexer.NewReindex(os.Getenv("REINDEXER"))
	err := DB.OpenNamespace(NSTelegramAccount, reindexer.DefaultNamespaceOptions(), TelegramAccount{})
	err = DB.OpenNamespace(NSDeviceRune, reindexer.DefaultNamespaceOptions(), DeviceRune{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ReindexerDB ok")

}
