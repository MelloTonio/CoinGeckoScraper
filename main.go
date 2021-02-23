package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Coin struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Percent1h  string `json:"Percent1h"`
	Percent24h string `json:"Percent24h"`
	Percent7d  string `json:"Percent7d"`
	MarketCap  string `json:"MarketCap"`
	Volume     string `json:"Volume"`
}

func IsLetter(str string) bool {
	for x := 0; x < len(str); x++ {
		ch := str[x]
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == ' ' || ch == '.' || ch == '0') {
			return false
		}
	}
	return true
}

func main() {
	var TopCoins []Coin
	//var nameArray []string

	CoinArr := make([][]string, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.coingecko.com"),
	)

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {

		tempCoinArr := make([]string, 0)

		element.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			// Pega todas informações (sem ser nome)
			row.ForEach("td > span", func(_ int, wantedText *colly.HTMLElement) {

				tempCoinArr = append(tempCoinArr, wantedText.DOM.Text())

			})

			// Busca apenas os nomes (a outra função não consegue pegar os nomes corretamente)
			row.ForEach("td", func(_ int, nameFinder *colly.HTMLElement) {
				var nameHelper string

				if IsLetter(nameFinder.Attr("data-sort")) {
					nameHelper = fmt.Sprintf("%s", nameFinder.Attr("data-sort"))
				}

				tempCoinArr = append(tempCoinArr, nameHelper)

			})

			CoinArr = append(CoinArr, tempCoinArr)

			tempCoinArr = []string{}

		})

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.coingecko.com/")

	for _, coin := range CoinArr {
		var TempCoin Coin

		TempCoin.Name = coin[8]
		TempCoin.Value = coin[0]
		TempCoin.Percent1h = coin[1]
		TempCoin.Percent24h = coin[2]
		TempCoin.Percent7d = coin[3]
		TempCoin.Volume = coin[4]
		TempCoin.MarketCap = coin[5]

		TopCoins = append(TopCoins, TempCoin)
	}

	writeJSON(TopCoins)

}

func writeJSON(data []Coin) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("Coins.json", file, 0644)
}
