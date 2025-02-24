package maxrevenue

func MaxRevenue(flightName_map map[string]float64)string{
	//flight with most revenue
	max1 := 0.0
	var max_flight string
	for key, value := range flightName_map {
		if value > max1 {
			max1 = float64(value)
			max_flight = key
		}
	}
	return max_flight
}