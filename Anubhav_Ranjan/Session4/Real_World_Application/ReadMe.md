You’ve been asked to build a tool for analyzing flight bookings for an airline company. The system should process booking data and provide useful information like total revenue per flight, the number of bookings in each seat class, and the most frequent flyers.



 Example

A list of bookings, where each booking is a string formatted like this:
"PassengerName:FlightNumber:SeatClass:Price"
Example: bookings := []string{
    "Alice:FL123:Economy:120.50",
    "Bob:FL123:Business:450.00",
    "Charlie:FL456:Economy:150.75",
    "Alice:FL456:Economy:150.75",
    "Bob:FL123:Economy:120.50",
}


 Tasks

Calculate total revenue for each flight.
Count how many bookings there are for each seat class (Economy, Business, First).
Figure out which passenger has the most bookings.
Find the flight that earned the most revenue.
Ignore invalid bookings.
Skip entries that are incorrectly formatted, have negative prices, or contain missing information.
Avoid counting duplicate bookings.
If the same passenger books the same flight and seat class more than once, it should only count once.


Edge Cases 

Bookings with empty strings or wrong formats.
Passengers booking multiple flights.
Flights with no valid bookings.


Expected  Output 

Revenue per flight: FL123: $690.50  
FL456: $301.50  
Bookings by seat class: Economy: 4  
Business: 1  
First: 0  
Passenger with the most bookings: Alice (2 bookings)  
Flight with the highest revenue: FL123 ($690.50)  
