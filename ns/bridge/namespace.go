package bridge

import (
	"github.com/gleba/radar-core/gates"
	"github.com/nats-io/go-nats"
	"github.com/restream/reindexer"
	"log"
)

type NS struct {
	Name string
}

var classMap = map[string]interface{}{}

func (ns NS) Register(class interface{}) {
	//println("domain namespace :", ns.Name)
	classMap[ns.Name] = class
}

func Add(name string) NS {
	return NS{
		Name: name,
	}
}

func (ns NS) On(f func(data []byte)) {
	log.Println("listen flow channel:", ns.Name)
	gates.Subscribe(ns.Name, func(msg *nats.Msg) {
		f(msg.Data)
	})
}
func (ns NS) Send(data interface{}) {
	gates.Publish(ns.Name, data)
}

//func (ns NS) GetById(id interface{}, target *interface{}) bool{
//	item, found := gates.Rei.
//		Query(ns.Name).
//		Where("id", reindexer.EQ, id).
//		Get()
//	if found {
//
//		return true
//	} else {
//		return false
//	}
//}

func (ns NS) GetByIdOrNew(id interface{}) (interface{}, bool) {
	item, found := gates.Rei.
		Query(ns.Name).
		Where("id", reindexer.EQ, id).
		Get()
	if found {
		return item, true
	} else {
		return nil, false
	}
}

func (ns NS) Save(data interface{}) {
	err := gates.Rei.Upsert(ns.Name, data)
	if err != nil {
		log.Print("ns.Save error in ", ns.Name)
		log.Fatal("↑", err)
	}
}

func (ns NS) SaveThis(data interface{}) {
	err := gates.Rei.Upsert(ns.Name, ns)
	if err != nil {
		log.Fatal(ns)
		log.Fatal("↑", err)
	}
}
