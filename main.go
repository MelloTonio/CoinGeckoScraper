package main

import (
	"github.com/mellotonio/coinfinder/src/Controller"
	Model "github.com/mellotonio/coinfinder/src/Models"
)

func main() {
	var options Model.Options

	// u (up) | d (down)
	options.Direction = "u"
	// -infinity ~ +infinity
	options.Percent = 1
	// h | d | 7d
	options.Time = "h"

	// Collect the coins then create a JSON with them
	CoinsJSON := Controller.FetchJson()

	// Filter the Coins by %
	Controller.FilterByPercent(CoinsJSON, options)

	Controller.GetAllCoins()
}
