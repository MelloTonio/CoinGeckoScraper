package main

import (
	"github.com/mellotonio/coinfinder/src/Controller"
	Model "github.com/mellotonio/coinfinder/src/Models"
)

func main() {
	var options Model.Options

	// u (up) | d (down) - Ex: If you want all coins that are '> 2%' -> direction:'u', percent = 2
	options.Direction = "d"
	// -infinity ~ +infinity
	options.Percent = -10
	// h | d | 7d
	options.Time = "7d"

	// Collect the coins then create a JSON with them
	CoinsJSON := Controller.FetchJson()

	// Filter the Coins by %
	Controller.FilterByPercent(CoinsJSON, options)

	Controller.GetAllCoins()
}
