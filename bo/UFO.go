package bo

import (
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ns"
	"github.com/gleba/radar-core/ux"
	"github.com/restream/reindexer"

	//"github.com/restream/reindexer"
	"time"
)

func init() {
	ns.UFO.Register(UFO{})
}

type UFO struct {
	CoinID     uint32 `reindex:"id,,pk"`
	CmcID      string
	NoiseClose float64
	NoiseStorm float64
	XVolumes   [3]float64
	AvgVolumes []float64
	FiRates    [3]float64
	FiScore    float64
	Score      float64
	Time       time.Time
}

func GetUfoPool() *[]UFO {
	query := gates.Rei.
		Query(ns.UFO.Name).
		Where("FiScore", reindexer.GT, 0.1).
		//Where("Time", reindexer.To, vars.MinCap).
		ReqTotal()

	iterator := query.Exec()
	defer iterator.Close()
	var items []UFO
	for iterator.Next() {
		items = append(items, *iterator.Object().(*UFO))
	}
	ux.Err(iterator.Error())
	return &items
}
