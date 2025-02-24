package maxbooking


func MaxBooking(name_map map[string]int)string{
	max := 0
	var max_name string
	for key, value := range name_map {
		if value > max {
			max = value
			max_name = key
		}
	}
	return max_name

}
