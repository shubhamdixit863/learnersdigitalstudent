package main

import (
	"fmt"
	"strconv"
	"strings"
)

func averageRevenue(bookings []string, name string) {
	var totalRevenue float64
	var count int

	for _, entry := range bookings {
		parts := strings.Split(entry, ":") 

		passengerName := parts[0]
		price, _ := strconv.ParseFloat(parts[3], 64) 
		if passengerName == name {
			totalRevenue += price
			count++
		}
	}
	average := totalRevenue / float64(count)
	fmt.Println( name, average)
}
func count(bookings []string,seat string)(int){
	var count int=0
	for _,entry:=range bookings{

		parts:=strings.Split(entry,":")

		class:=parts[2]
		if(class==seat){
			count++
		}
	}
	return count
}
func mostbooking(bookings []string){
	counts := make(map[string]int)
	var maxName string
	maxCount := 0

	for _, entry := range bookings {
		name := strings.Split(entry, ":")[0]
		counts[name]++
		if counts[name] > maxCount {
			maxCount = counts[name]
			maxName = name
		}
	}
	fmt.Println(maxName,maxCount)

}

func maxrevenue(bookings []string){
	revenue := make(map[string]float64)
	var maxFlight string
	maxRevenue := 0.0

	for _, entry := range bookings {
		parts := strings.Split(entry, ":")
		flight := parts[1]
		price, _ := strconv.ParseFloat(parts[3], 64)
		revenue[flight] += price
		if revenue[flight] > maxRevenue {
			maxRevenue = revenue[flight]
			maxFlight = flight
		}
	}
	fmt.Println(maxFlight,maxRevenue)
}
func main() {
	bookings := []string{
		"Alice:FL123:Economy:120.50",
		"Bob:FL123:Business:450.00",
		"Charlie:FL456:Economy:150.75",
		"Alice:FL456:Economy:150.75",
		"Bob:FL123:Economy:120.50",
	}

	var name string
	fmt.Println("Enter passenger name for average revenue calculation:")
	fmt.Scanln(&name)
	averageRevenue(bookings, name)
	var seat string
	fmt.Println("Enter Seat Class")
	fmt.Scanln(&seat)
	num:=count(bookings,seat)
	fmt.Println(num)
	mostbooking(bookings)
	maxrevenue(bookings)
}
