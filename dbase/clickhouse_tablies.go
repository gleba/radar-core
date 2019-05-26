package dbase

import "fmt"

const (
	TableCmcPulse    = "CmcPulse"
	TableCmcDaysPulse = "CmcDaysPulse"
)

var tablePulseVol = `	
	VolumeUSD  Float64,
	VolumeBTC  Float64,
	PriceUSD   Float64,
	PriceBTC   Float64,`

var tableMarketCap = `
	MarketCapUSD Float64,
	MarketCapBTC Float64,`

var tablies = []string{
	fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	%s
	%s	
	CoinID      UInt32,
	Time  	  	DateTime
)
ENGINE=MergeTree()
ORDER BY (CoinID, Time)`, TableCmcPulse, tablePulseVol, tableMarketCap),
fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	Open      Float64,
	Close     Float64,
	High      Float64,
	Low       Float64,
	Volume    Float64,
	MarketCap Float64,
	CoinID    UInt32,
	Date  		Date
)
ENGINE=MergeTree()
ORDER BY (CoinID, Date)
`, TableCmcDaysPulse),
	//fmt.Sprintf(`
	//CREATE TABLE IF NOT EXISTS CmcPulseByMarkets (
	//	%s
	//	%s
	//	CoinId      UInt32,
	//	MarketId    UInt32,
	//	Time  		DateTime
	//)
	//engine=MergeTree(Time, (CoinId), 8192)
	//`, t✓blePulseVol, tableMarketCap),
}
