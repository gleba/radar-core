package bo

import (
	"database/sql"
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ux"
	"log"
	"time"
)

type PriceVol struct {
	VolumeUSD float64
	VolumeBTC float64
	PriceUSD  float64
	PriceBTC  float64
}
type MarketCap struct {
	MarketCapUSD float64
	MarketCapBTC float64
}

type CoinPulse struct {
	PriceVol
	MarketCap
	CoinId uint32
	Time   time.Time
}

//----- next SQL Writer ↓↓↓

type PulseWriter struct {
	tx        *sql.Tx
	stmt      *sql.Stmt
	hasUpdate bool
	Count     int
}

func CreatePulseWriter() *PulseWriter {
	writer := PulseWriter{
		Count: 0,
	}
	var err error
	writer.tx, err = gates.SqlX.Begin()
	ux.Err(err)
	writer.stmt, err = writer.tx.Prepare("INSERT INTO CmcPulse (CoinId, Time, VolumeUSD, VolumeBTC, MarketCapUSD, MarketCapBTC,PriceUSD,PriceBTC) VALUES (?,?,?,?,?,?,?,?)")
	ux.Err(err)
	return &writer
}

func (self *PulseWriter) Add(pulse CoinPulse) {
	_, err := self.stmt.Exec(
		pulse.CoinId, pulse.Time,
		pulse.VolumeUSD, pulse.VolumeBTC, pulse.MarketCapUSD, pulse.MarketCapBTC, pulse.PriceUSD, pulse.PriceBTC)
	ux.Err(err)
	self.Count = 1 + self.Count
}
func (w *PulseWriter) Commit() {
	if w.Count >= 1 {
		ux.Err(w.tx.Commit())
		log.Println("pulse change ", w.Count, "elements")
		w.Count = 0
	}
}
