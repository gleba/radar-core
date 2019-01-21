package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/restream/reindexer"
)

func init() {
	ns.CoinRates.Register(CoinRates{})
}

const (
	StateUnknown uint = iota
	StateDetected
	StateTargeted
	StateComplicated
)

type CoinRates struct {
	CoinID           uint32 `reindex:"id,,pk"`
	State            uint   `reindex:"state"`
	VolatilityRate   float64
	VolatilityRateRT float64
	ImpulseRate      float64
	AvgPrice         float64
	Noise            float64
}

func GetCoinRatesByState(state uint) []CoinRates {
	var query *reindexer.Query
	//if len(ak) > 0 {
	query = gates.Rei.
		Query(ns.CoinRates.Name).
		//Where("state", reindexer.EQ, state).
		ReqTotal()
	iterator := query.Exec()
	defer iterator.Close()
	var rates []CoinRates
	for iterator.Next() {
		rates = append(rates, *iterator.Object().(*CoinRates))
	}
	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}
	return rates
}
