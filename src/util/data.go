package util

// TickerResponse API response from CoinMarketCap
type TickerResponse []TickerListing

// TickerListing individual coin data from response
type TickerListing struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Rank            string `json:"rank"`
	PriceUSD        string `json:"price_usd"`
	PriceBTC        string `json:"price_btc"`
	VolUSD24H       string `json:"24h_volume_usd"`
	MarketCapUSD    string `json:"market_cap_usd"`
	AvailableSupply string `json:"available_supply"`
	TotalSupply     string `json:"total_supply"`
	MaxSupply       string `json:"max_supply"`
	PctChange1H     string `json:"percent_change_1h"`
	PctChange24H    string `json:"percent_change_24h"`
	PctChange7d     string `json:"percent_change_7d"`
	LastUpdated     string `json:"last_updated"`
}

/*
{
        "id": "ethereum",
        "name": "Ethereum",
        "symbol": "ETH",
        "rank": "2",
        "price_usd": "717.405",
        "price_btc": "0.0494806",
        "24h_volume_usd": "2496830000.0",
        "market_cap_usd": "69250817665.0",
        "available_supply": "96529600.0",
        "total_supply": "96529600.0",
        "max_supply": null,
        "percent_change_1h": "0.89",
        "percent_change_24h": "2.9",
        "percent_change_7d": "1.52",
        "last_updated": "1514078656"
    }
*/
