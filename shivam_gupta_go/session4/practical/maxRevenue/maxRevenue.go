package maxrevenue

import "strconv"

func MaxRevenue(a [][]string) (string, float64) {
	mp := make(map[string]float64)

	for _, num := range a {
		flightNo := num[1]
		price, _ := strconv.ParseFloat(num[3], 64)
		mp[flightNo] += price
	}
	flightName := ""
	var revenue float64
	for key, value := range mp {
		if value > revenue {
			revenue = value
			flightName = key
		}
	}
 return flightName,revenue
}
