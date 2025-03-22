# 🚗 Vechile Rental System

The Vehicle Rental System is a Go-based program designed to manage vehicle rentals for a company that offers Cars, Bikes, and Trucks. The system calculates rental costs dynamically based on the hours rented.

## ⛏️ Features

1. Vehicle Management: Supports Cars, Bikes, and Trucks.

2. Rental Cost Calculation: Computes cost dynamically based on hourly rates.

3. Error Handling: Includes proper validation and error messages.

## 💻 Installation and setup

1. Extract the downloaded zip folder
2. Find a project struture containing the overall division of two folder {"cmd"} and {"internal"}
3. The {"internal"} folder contains all the files and folders required for the proper execution of the program
4. The {"cmd"} folder contains the main.go file which is the entry point of the program.
5. To check if the program is correctly installed run the command -> {"./binary.exe"}

## 📤 Example output

Vehicle: Car | Model: Toyota | Cost: $50.00<br />
Vehicle: Car | Model: Honda | Cost: $75.00<br />
Vehicle: Bike | Model: Yamaha | Cost: $25.00<br />
Vehicle: Bike | Model: Suzuki | Cost: $15.00<br />
Vehicle: Truck | Model: Isuzu | Cost: $100.00<br />
Vehicle: Truck | Model: Hino | Cost: $125.00<br />
2025/03/21 12:41:09 vehicle must have a name<br />
2025/03/21 12:41:09 rental rate must be greater than zero<br />

## ⚠️ Error Handeling

1. ErrInvalidHours: rental hours must be greater than zero
2. ErrNilVehicle: vehicle cannot be nil
3. ErrUnNamedVehicle: vehicle must have a name
4. ErrInvalidRentalRate: rental rate must be greater than zero
