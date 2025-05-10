package totalRevenue

import "strconv"

func TotalRevenue(a [][]string) map[string]float64 {
	mp := make(map[string]float64)

	for _, num := range a {
		flightNo := num[1]
		price,_ := strconv.ParseFloat(num[3],64)
		mp[flightNo]+=price
	}
	return mp
}