package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mellotonio/coinfinder/src/Controller"
	"github.com/mellotonio/coinfinder/src/Models"
)

func AllRoutes(r *gin.Engine) {
	//  ex: {{baseurl}}/all?pageLimit=3
	r.GET("/all", func(c *gin.Context) {
		pageLimitStr := c.Query("pageLimit")

		pageLimitInt, _ := strconv.Atoi(pageLimitStr)

		// Update Coins
		Controller.GetAllCoins(pageLimitInt)

		CoinsJSON := Controller.FetchJson("Coins.json")

		key := fmt.Sprintf("%d_pages", pageLimitInt)

		c.JSON(200, gin.H{
			key: CoinsJSON,
		})
	})

	// ex: {{baseurl}}/filter?pageLimit=5&direction=down&percent=100&time=7d
	r.GET("/filter", func(c *gin.Context) {
		pageLimitStr := c.Query("pageLimit")
		directionStr := c.Query("direction")
		percentStr := c.Query("percent")
		timeStr := c.Query("time")

		pageLimitInt, _ := strconv.Atoi(pageLimitStr)

		var options Models.Options

		if percentFloat, err := strconv.ParseFloat(percentStr, 32); err == nil {
			// -infinity ~ +infinity
			options.Percent = percentFloat
		}

		// u (up) | d (down) - Ex: If you want all coins that are '> 2%' -> direction:'u', percent = 2
		options.Direction = directionStr
		// h | d | 7d
		options.Time = timeStr

		// Update Coins
		Controller.GetAllCoins(pageLimitInt)

		JsonName := fmt.Sprintf("src/FilteredCoins/%s/Coins_%s_%d%%_%s.json", options.Time, options.Direction, int(options.Percent), options.Time)

		ParseAllCoins := Controller.FetchJson("Coins.json")

		Controller.FilterByPercent(ParseAllCoins, options)

		ParseFilteredCoins := Controller.FetchJson(JsonName)

		c.JSON(200, gin.H{
			"filtered_coins": ParseFilteredCoins,
		})
	})
}
