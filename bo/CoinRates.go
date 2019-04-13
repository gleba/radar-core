package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/restream/reindexer"
)

func init() {
	ns.CoinRates.Register(CoinRate{})
}

const (
	StateUnknown uint = iota
	StateDetected
	StateTargeted
	StateComplicated
)

type CoinRate struct {
	//CoinID           uint32 `reindex:"id,,pk"`

	VolatilityRate   float64
	VolatilityRateRT float64
	ImpulseRate      float64
	Noise            float64
	CloseChan        string
	//Time             time.Time
	//Valid            bool
}

func GetCoinRatesByState(state uint) []CoinRate {
	var query *reindexer.Query
	query = gates.Rei.
		Query(ns.CoinRates.Name).
		ReqTotal()
	iterator := query.Exec()
	defer iterator.Close()
	var rates []CoinRate
	for iterator.Next() {
		rates = append(rates, *iterator.Object().(*CoinRate))
	}
	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}
	return rates
}
