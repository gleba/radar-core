package bo

import (
	"github.com/gleba/radar-core/ns"
	//"github.com/restream/reindexer"
	"time"
)

func init() {
	ns.UFO.Register(UFO{})
}

type UFO struct {
	CoinID uint32 `reindex:"id,,pk"`
	//CmcID      string
	NoiseClose float64
	NoiseStorm float64
	XVolumes   [3]float64
	AvgVolumes []float64
	AvgPrices  []float64
	ARates     [3]float64
	//PriseRate   float64
	//ImpulseRate float64
	//Score       float64
	Time time.Time
}
