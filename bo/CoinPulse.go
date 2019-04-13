package bo

import (
	"database/sql"
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ux"
	"github.com/gleba/radar-core/vars"
	"log"
	"time"
)

type PriceVol struct {
	VolumeUSD float64 `db:"VolumeUSD"`
	VolumeBTC float64 `db:"VolumeBTC"`
	PriceUSD  float64 `db:"PriceUSD"`
	PriceBTC  float64 `db:"PriceBTC"`
}
type MarketCap struct {
	MarketCapUSD float64 `db:"MarketCapUSD"`
	MarketCapBTC float64 `db:"MarketCapBTC"`
}

type CoinPulse struct {
	PriceVol
	MarketCap
	CoinId uint32    `db:"CoinId"`
	Time   time.Time `db:"Time"`
}

func (pulse *CoinPulse) IsAlive() bool {
	return pulse.VolumeUSD > vars.MinVolume && pulse.MarketCapUSD > vars.MinCap
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
