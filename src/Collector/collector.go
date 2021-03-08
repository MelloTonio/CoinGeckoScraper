package Collector

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/mellotonio/coinfinder/src/Helpers"
	Model "github.com/mellotonio/coinfinder/src/Models"
)

func GetCoinGecko(pageLimit int) []Model.Coin {
	var TopCoins []Model.Coin
	coinArray := make([][]string, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.coingecko.com"),
	)

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {

		// Cria um array temporário de moedas
		tempCoinArr := []string{"", "", "", "", "", "", ""}
		//pointer := 0

		element.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			// Pega todas informações (sem ser nome)
			row.ForEach("td > span", func(pos int, wantedText *colly.HTMLElement) {
				tempCoinArr[pos] = wantedText.Text
				//	tempCoinArr = append(tempCoinArr, wantedText.Text)
			})

			// Busca apenas os nomes (a outra função não consegue pegar os nomes corretamente)
			row.ForEach("td", func(_ int, nameFinder *colly.HTMLElement) {
				var nameHelper string

				if Helpers.IsLetter(nameFinder.Attr("data-sort")) {
					nameHelper = fmt.Sprintf("%s", nameFinder.Attr("data-sort"))
				}

				// Evita que espaços em branco sejam tratados como validos
				if nameHelper != "" && len(nameHelper) > 2 {
					tempCoinArr[6] = nameHelper
					//tempCoinArr = append(tempCoinArr, nameHelper)
				}

			})

			// Verifica se algum valor não existe, o colly pula rows quando elas não existem.
			if !Helpers.Contains(tempCoinArr, "") {
				// Adiciona uma moeda ao array de moedas
				coinArray = append(coinArray, tempCoinArr)
			}

			// Zera o array de moedas temporárias
			tempCoinArr = []string{"", "", "", "", "", "", ""}

		})

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.coingecko.com/")
	for pageNumber := 1; pageNumber < pageLimit; pageNumber++ {
		page := fmt.Sprintf("https://www.coingecko.com/pt?page=%d", pageNumber)

		collector.Visit(page)
	}

	// Cria uma Moeda para cada moeda recebida do CoinGecko
	for _, coin := range coinArray {
		var TempCoin Model.Coin

		fmt.Println(coin)

		TempCoin.Value = coin[0]
		TempCoin.Percent1h = coin[1]
		TempCoin.Percent24h = coin[2]
		TempCoin.Percent7d = coin[3]
		TempCoin.Volume = coin[4]
		TempCoin.MarketCap = coin[5]
		TempCoin.Name = coin[6]

		TempCoin.Date = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		TopCoins = append(TopCoins, TempCoin)
	}

	return TopCoins
}
