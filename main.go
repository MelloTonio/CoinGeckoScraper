package main

import (
	"github.com/gin-gonic/gin"
	Model "github.com/mellotonio/coinfinder/src/Models"
	"github.com/mellotonio/coinfinder/src/routes"
)

func main() {
	r := gin.Default()

	routes.AllRoutes(r)

	r.Run()

	var options Model.Options

	// u (up) | d (down) - Ex: If you want all coins that are '> 2%' -> direction:'u', percent = 2
	options.Direction = "d"
	// -infinity ~ +infinity
	options.Percent = 100
	// h | d | 7d
	options.Time = "7d"

	// Collect the coins then create a JSON with them
	//CoinsJSON := Controller.FetchJson()

	// Filter the Coins by %
	//Controller.FilterByPercent(CoinsJSON, options)

	//Controller.GetAllCoins()
}
