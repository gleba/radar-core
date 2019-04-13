package bo

import (
	"database/sql"
	"github.com/gleba/radar-core/gates"
	"github.com/gleba/radar-core/ux"
	"sync"
	"time"
)

type HistoryPulse struct {
	Open      float64   `db:"Open"`
	Close     float64   `db:"Close"`
	High      float64   `db:"High"`
	Low       float64   `db:"Low"`
	Volume    float64   `db:"Volume"`
	MarketCap float64   `db:"MarketCap"`
	CoinId    uint32    `db:"CoinId"`
	Date      time.Time `db:"Date"`
}

type HistoryPulseWriter struct {
	tx    *sql.Tx
	stmt  *sql.Stmt
	Count int
}

var hack sync.Map

func GetHistoryToDate(id uint32, date time.Time) []HistoryPulse {
	var items []HistoryPulse
	v, f := hack.Load(id)
	if f {
		items = v.([]HistoryPulse)
	} else {
		ux.Err(gates.SqlX.Select(&items,
			"select Open, Volume, Close, High, Low, MarketCap, MarketCap, Date from CmcDayPulse where CoinId = ? order by Date DESC ", id))
		hack.Store(id, items)
	}
	var fitems []HistoryPulse
	for _, i := range items {
		if i.Date.Unix() <= date.Unix() {
			fitems = append(fitems, i)
		}
	}
	return fitems
}

func GetHistory(id uint32, limit int) []HistoryPulse {
	var items []HistoryPulse
	ux.Err(gates.SqlX.Select(&items,
		"select Open, Volume, Close, High, Low, MarketCap, MarketCap, Date from CmcDayPulse where CoinId = ? order by Date DESC limit ?", id, limit))
	return items
}

func CreateHistoryPulseWriter() *HistoryPulseWriter {
	writer := HistoryPulseWriter{
		Count: 0,
	}
	var err error
	writer.tx, err = gates.SqlX.Begin()
	ux.Err(err)
	writer.stmt, err = writer.tx.Prepare(
		"INSERT INTO CmcDayPulse (CoinId, Date, Open, Close, High, Low, Volume, MarketCap ) VALUES (?,?,?,?,?,?,?,?)")
	ux.Err(err)
	return &writer
}

//
func (w *HistoryPulseWriter) Add(pulse HistoryPulse) {
	_, err := w.stmt.Exec(
		pulse.CoinId, pulse.Date,
		pulse.Open, pulse.Close, pulse.High, pulse.Low, pulse.Volume, pulse.MarketCap)
	ux.Err(err)
	w.Count = 1 + w.Count
}
func (w *HistoryPulseWriter) Commit() {
	if w.Count >= 1 {
		ux.Err(w.tx.Commit())
		w.Count = 0
	}
}
