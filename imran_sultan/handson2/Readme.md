This is a simple Vehicle Rental System implemented in Go. It allows users to:

Add different types of vehicles (Car, Bike, Truck)

List available vehicles with their hourly rental rates

Calculate rental costs for a specific vehicle

vehicle/

This directory contains vehicle-related files:

vehicle.go: Defines the Vehicle interface and BaseVehicle struct.

car.go: Defines the Car struct.

bike.go: Defines the Bike struct.

truck.go: Defines the Truck struct.

rental/

This directory contains the rental management logic:

rental_manager.go: Manages vehicle rentals, adding vehicles, and calculating rental costs.


Run the Application:

go run main.go




Example Output:

Car (ID: C123) - Rate: 10.00/hour
Bike (ID: B456) - Rate: 5.00/hour
Truck (ID: T789) - Rate: 20.00/hour
Rental Cost for Car (5 hours): 50.00