package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mellotonio/coinfinder/src/Collector"
	Colors "github.com/mellotonio/coinfinder/src/Color"
	"github.com/mellotonio/coinfinder/src/Helpers"
	Model "github.com/mellotonio/coinfinder/src/Models"
)

func FetchJson() []Model.Coin {
	Top100Coins := Collector.GetCoinGecko()
	Helpers.WriteJSON(Top100Coins, "Coins.json")

	jsonFile, err := os.Open("Coins.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Coin.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var Coins []Model.Coin

	json.Unmarshal(byteValue, &Coins)

	return Coins

}

// Filter the Coin list by (up or down) %
// Options Ex: 'Options: { Direction: "d", Percent: 5}'

func FilterByPercent(CoinsList []Model.Coin, options Model.Options) string {
	color.Yellow("Time Range: %s - Percentage: %d%% - Direction: %s", options.Time, int(options.Percent), options.Direction)

	var FilteredCoins []Model.Coin
	var removePercent string

	if options.Direction != "d" && options.Direction != "u" {
		color.Red("Invalid 'Percent Direction'. Please use: u (up) or d (down)")
		return ""
	}

	for _, coin := range CoinsList {
		var coinPercentage string

		switch options.Time {

		case "h":
			coinPercentage = coin.Percent1h
			removePercent = strings.Replace(coin.Percent1h, "%", "", 1)
			break

		case "d":
			coinPercentage = coin.Percent24h
			removePercent = strings.Replace(coin.Percent24h, "%", "", 1)
			break

		case "7d":
			coinPercentage = coin.Percent7d
			removePercent = strings.Replace(coin.Percent7d, "%", "", 1)
			break

		default:
			fmt.Println("Invalid time range!")
			return ""
		}

		if coinPercent, err := strconv.ParseFloat(removePercent, 16); err == nil {

			switch options.Direction {

			case "d":
				if coinPercent < options.Percent {
					coin.Description = fmt.Sprintf("%s decreased %.2f%% in a '%s' period - DATE: %s", coin.Name, coinPercent, options.Time, coin.Date)

					FilteredCoins = append(FilteredCoins, coin)

					color.Blue("Coin: %s - Drop Percentage: %s\n", coin.Name, coinPercentage)
				}
				break

			case "u":
				if coinPercent > options.Percent {
					coin.Description = fmt.Sprintf("%s increased %f in a '%s' period - DATE: %s", coin.Name, coinPercent, options.Time, coin.Date)

					FilteredCoins = append(FilteredCoins, coin)

					color.Blue(Colors.White+"Coin: %s - Up Percentage: %s\n", coin.Name, coinPercentage)
				}
				break
			}

		}

	}

	// It will create a file -> Coins_(up or down)_(Percent up or down)_(Time range).json
	JsonName := fmt.Sprintf("src/FilteredCoins/%s/Coins_%s_%d%%_%s.json", options.Time, options.Direction, int(options.Percent), options.Time)
	Helpers.WriteJSON(FilteredCoins, JsonName)

	color.Yellow("Create a new file in: %s", JsonName)

	return "Success"
}

func GetAllCoins() {
	Top100Coins := Collector.GetCoinGecko()
	Helpers.WriteJSON(Top100Coins, "Coins.json")

	color.Yellow("Updated 'Coins.json' successfully!")
}
