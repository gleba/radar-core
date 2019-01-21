package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/gleba/radar-core/ux"
	//"github.com/restream/reindexer"
	"time"
)

func init() {
	ns.UFA.Register(UFA{})
}

type UFA struct {
	CoinID      uint32 `reindex:"id,,pk"`
	Define      string
	NoiseRate   float64
	ImpulseRate float64
	TimeDetect  time.Time
	TimeMax     time.Time
	TimeEnd     time.Time
}

func GetUfaPool() *[]UFA {
	query := gates.Rei.
		Query(ns.UFA.Name).
		//Where("FiScore", reindexer.GT, 0.1).
		//Where("Time", reindexer.To, vars.MinCap).
		ReqTotal()

	iterator := query.Exec()
	defer iterator.Close()
	var items []UFA
	for iterator.Next() {
		items = append(items, *iterator.Object().(*UFA))
	}
	ux.Err(iterator.Error())
	return &items
}
