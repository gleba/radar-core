package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/restream/reindexer"
)

type Coin struct {
	Id     uint32 `reindex:"id,,pk"`
	CmcId  string `reindex:"cmc"`
	Name   string
	Symbol string
	Rate   float64
	Pulse  CoinPulseOld
}

func init() {
	ns.Coin.Register(Coin{})
}

func GetCoinsPool() []Coin {
	var query *reindexer.Query
	//if len(ak) > 0 {
	query = gates.Rei.
		Query(ns.Coin.Name).
		ReqTotal()
	//} else {
	//	query = gates.Rei.
	//		Query(ns.Coin.Name).
	//		Where("Pulse.VolumeUSD", reindexer.GT, vars.MinVolume).
	//		Where("Pulse.MarketCapUSD", reindexer.GT, vars.MinCap).
	//		ReqTotal()
	//}
	iterator := query.Exec()
	defer iterator.Close()
	var coins []Coin
	for iterator.Next() {
		coins = append(coins, *iterator.Object().(*Coin))
	}
	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}
	return coins
}

//
//func GetCoinByID(id uint32) (Coin, bool) {
//	//item, found := gates.Rei.
//	//	Query(ns.Coin.Name).
//	//	Where("id", reindexer.EQ, id).
//	//	Get()
//	//if found {
//	//	coin := item.(*Coin)
//	//	return *coin, found
//	//} else {
//	//	return Coin{
//	//		Id: id,
//	//	}, found
//	//}
//}
//
//func (item Coin) Save() {
//	//err := gates.Rei.Upsert(ns.Coin.Name, item)
//	//if err != nil {
//	//	log.Fatal(item)
//	//	log.Fatal("↑", err)
//	//}
//}
