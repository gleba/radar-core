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

type CoinPulseOld struct {
	PriceVol
	MarketCap
	CoinId uint32    `db:"CoinId"`
	Time   time.Time `db:"Time"`
}

type CoinPulse struct {
	PriceVol
	MarketCap
	CoinID uint32    `db:"CoinID"`
	Time   time.Time `db:"Time"`
}
//type CoinPulse struct {
//	PriceVol
//	MarketCap
//	CoinID uint32    `db:"CoinID"`
//	Time   time.Time `db:"Time"`
//}

func (pulse *CoinPulseOld) IsAlive() bool {
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
	ux.Safe(err)
	writer.stmt, err = writer.tx.Prepare("INSERT INTO CmcPulse (CoinId, Time, VolumeUSD, VolumeBTC, MarketCapUSD, MarketCapBTC,PriceUSD,PriceBTC) VALUES (?,?,?,?,?,?,?,?)")
	ux.Safe(err)
	return &writer
}

func (self *PulseWriter) Add(pulse CoinPulseOld) {
	_, err := self.stmt.Exec(
		pulse.CoinId, pulse.Time,
		pulse.VolumeUSD, pulse.VolumeBTC, pulse.MarketCapUSD, pulse.MarketCapBTC, pulse.PriceUSD, pulse.PriceBTC)
	ux.Safe(err)
	self.Count = 1 + self.Count
}
func (w *PulseWriter) Commit() {
	if w.Count >= 1 {
		ux.Safe(w.tx.Commit())
		log.Println("pulse change ", w.Count, "elements")
		w.Count = 0
	}
}
