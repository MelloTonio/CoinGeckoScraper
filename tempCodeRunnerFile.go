package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Coin struct {
	Zap        string `json:"zao"`
	Zap2       string `json:"zao2"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Percent24h string `json:"Percent24h"`
	Percent7d  string `json:"Percent7d"`
	MarketCap  string `json:"MarketCap"`
	Volume     string `json:"Volume"`
}

func main() {
	//	allFacts := make([]Coin, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.coingecko.com"),
	)

	collector.OnHTML("tbody > tr", func(element *colly.HTMLElement) {

		//CoinArr := make([]string, 0)

		/*		element.ForEach("td > span", func(_ int, el *colly.HTMLElement) {
				CoinArr = append(CoinArr, el.DOM.Text())
			})*/

		fmt.Println(element.DOM.Text())
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.coingecko.com/")

	//writeJSON(allFacts)
}

/*func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}*/
