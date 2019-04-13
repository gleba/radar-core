package flow

import (
	"github.com/gleba/radar-core/bo"
	"github.com/gleba/radar-core/ns"
	"github.com/gleba/radar-core/ux"
	"github.com/vmihailenco/msgpack"
)

func OnCoinPulse(f func(*[]bo.CoinPulse)) {
	ns.CoinsPulse.On(func(data []byte) {
		var v []bo.CoinPulse
		ux.Err(msgpack.Unmarshal(data, &v))
		f(&v)
	})
}
