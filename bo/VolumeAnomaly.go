package bo

import (
	"github.com/gleba/radar-core/ns"
	//"github.com/restream/reindexer"
	"time"
)

func init() {
	ns.VolumeAnomaly.Register(VolumeAnomaly{})
}

type VolumeAnomaly struct {
	CoinID     uint32 `reindex:"id,,pk"`
	VolumeRate float64
	TimeDetect time.Time
	TimeMax    time.Time
	TimeEnd    time.Time
}
