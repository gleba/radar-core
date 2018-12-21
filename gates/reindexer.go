package gates

import (
	"fmt"
	"github.com/restream/reindexer"
	"log"
	"os"
)

type ReiClasses struct {
	Name  string
	Class interface{}
}

var Rei *reindexer.Reindexer

//const NSTelegramAccount string = "telegram_account"
//type NameSpace struct {
//	name string
//	class interface{}
//}
func OpenReIndexer(nameSpaces []ReiClasses) {
	Rei = reindexer.NewReindex(os.Getenv("REINDEXER"))
	for _, nameSpace := range nameSpaces {
		//fmt.Println("OpenReIndexer")
		log.Println("open ReIndexer namespace :", nameSpace.Name)
		err := Rei.OpenNamespace(nameSpace.Name, reindexer.DefaultNamespaceOptions(), nameSpace.Class)
		if err != nil {
			log.Fatal(err)
		}
	}
	//log.Println("open gate: ReIndexer")
}

func noErr(err error, item interface{}) {
	if err != nil {
		fmt.Println(item)
		log.Fatal("↑", err)
	}
}
