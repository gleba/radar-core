package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/restream/reindexer"
	"time"
)

type CoinMarkets struct {
	CoinID  uint32 `reindex:"id,,pk"`
	Markets []CoinMarket
}

type CoinMarket struct {
	Id    uint32
	Name  string
	Pairs []MarketPair
}

type MarketPair struct {
	PriceVol
	Name       string
	Url        string
	UpdateTime time.Time
}

func init() {
	ns.CoinMarkets.Register(CoinMarkets{})
}

func GetMarketForCoinID(id uint32) (CoinMarkets, bool) {
	item, found := gates.Rei.
		Query(ns.CoinMarkets.Name).
		Where("id", reindexer.EQ, id).
		Get()
	if found {
		return *item.(*CoinMarkets), found
	} else {
		return CoinMarkets{}, found
	}
}
