package Models

type Coin struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Percent1h   string `json:"Percent1h"`
	Percent24h  string `json:"Percent24h"`
	Percent7d   string `json:"Percent7d"`
	MarketCap   string `json:"MarketCap"`
	Volume      string `json:"Volume"`
	Date        string `json:"Date"`
	Description string `json:"Description"`
}
