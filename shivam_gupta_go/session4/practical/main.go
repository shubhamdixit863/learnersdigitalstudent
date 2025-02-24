package main

import (
	"bufio"
	"fmt"
	"os"
	"practical/classBooking"
	 "practical/maxrevenue"
	"practical/passenger"
	"practical/totalRevenue"
	"strings"
)

var bookings []string


func main() {

	fmt.Println("Enter flight bookings (Format: Name:FlightNo:Class:Price). Type 'done' to stop:")
	for {
		fmt.Println("Enter bookings : ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "done" {
			break
		}
		bookings = append(bookings, input)
	}

	var parsedBookings [][]string

	for _, booking := range bookings {
		value := strings.Split(booking, ":")
		parsedBookings = append(parsedBookings, value)

	}
	// fmt.Println(parsedBookings)

	totalRevenue1 := totalRevenue.TotalRevenue(parsedBookings)
	fmt.Println("total revenue of each flight")
	fmt.Println(totalRevenue1)
	fmt.Println("bookings with class")
 class1:=classbooking.Classbooking(parsedBookings)
 fmt.Println(class1)

 name,rev:=maxrevenue.MaxRevenue(parsedBookings)
 fmt.Println("max revenue flight name : ",name," revenue ",rev)

 
pass1,book:=passenger.Passenger(parsedBookings)

fmt.Println("Passenger with most bookings :",pass1,book)
}
