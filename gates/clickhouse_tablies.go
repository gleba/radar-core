package gates

import "fmt"

const (
	TableCmcPulse    = "CmcPulse"
	TableCmcDayPulse = "CmcDayPulse"
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
	CoinId      UInt32,
	Date  		Date MATERIALIZED toDate(Time),
	Time  		DateTime
) 
engine=MergeTree(Date, Time, 8192)
`, TableCmcPulse, tablePulseVol, tableMarketCap),

	fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	Open      Float64,
	Close     Float64,
	High      Float64,
	Low       Float64,
	Volume    Float64,
	MarketCap Float64,
	CoinId    UInt32,
	Date  		Date
)
engine=ReplacingMergeTree()
PARTITION BY (Date) 
ORDER BY (CoinId, Date)
`, TableCmcDayPulse),
	//fmt.Sprintf(`
	//CREATE TABLE IF NOT EXISTS CmcPulseByMarkets (
	//	%s
	//	%s
	//	CoinId      UInt32,
	//	MarketId    UInt32,
	//	Time  		DateTime
	//)
	//engine=MergeTree(Time, (CoinId), 8192)
	//`, tablePulseVol, tableMarketCap),
}
