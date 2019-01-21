package bridge

import (
	"fmt"
	"github.com/gleba/radar-core/gates"
)

//type TypeUFO = *[]store.CoinPulse
//type Rei struct {
//	//ns.NS
//}
//var classMap = map[string]interface{}{}
//
//func Add(namespace bridge.NS, class interface{}) bool{
//	classMap[namespace.Name] = class
//	return true
//}

func ToReIndexer(namespaces ...NS) {
	var classes []gates.ReiClasses
	for _, n := range namespaces {
		v := classMap[n.Name]
		if v == nil {
			fmt.Println(n)
			panic("ToReIndexer: class not found in classMap → " + n.Name)
		}
		classes = append(classes, gates.ReiClasses{
			Name:  n.Name,
			Class: v,
		})
	}
	gates.OpenReIndexer(classes)
}

//func Name()  {
//
//}
//func (ns NS) On(f func(data []byte)) {
//	log.Println("listen flow channel:", ns.name)
//	gates.Subscribe(ns.name, func(msg *nats.Msg) {
//		f(msg.Data)
//	})
//}
//func (ns NS) Send(data interface{}) {
//	gates.Publish(ns.name, data)
//}
//
//func (ns NS) GetByIdOrNew(id interface{}) (interface{},) {
//	item, found := gates.Rei.
//		Query(ns.name).
//		Where("id", reindexer.EQ, id).
//		Get()
//	if found {
//		return item
//	} else {
//		return ns.class
//	}
//}
//
//func (ns Rei) Name() string{
//	return "xx"
//}
//func (ns NS) Save(data interface{}) {
//	err := gates.Rei.Upsert(ns.name, data)
//	if err != nil {
//		log.Fatal(ns)
//		log.Fatal("↑", err)
//	}
//}
//
//func (ns NS) SaveThis(data interface{}) {
//	err := gates.Rei.Upsert(ns.name, ns)
//	if err != nil {
//		log.Fatal(ns)
//		log.Fatal("↑", err)
//	}
//}
//
