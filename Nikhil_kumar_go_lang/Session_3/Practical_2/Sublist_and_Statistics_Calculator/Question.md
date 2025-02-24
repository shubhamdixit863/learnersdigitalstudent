Problem Statement: Build a Sublist and Statistics Calculator Using the Slicing Operator in Golang

Overview:

Create a CLI application in Golang that performs various operations on a list of integers provided by the user. The application should use the slicing operator (:) to create sublists and calculate statistics like sum, average, and maximum of the sliced portion.

Requirements:
	1.	Input List of Integers:
	•	Prompt the user to enter a list of integers separated by spaces (e.g., 10 20 30 40 50).
	•	Store these integers in a slice.
	2.	Slice the List:
	•	Ask the user to input start and end indices.
	•	Use the slice operator to create a sublist from the original list.
	•	Validate that indices are within bounds and that start <= end.
	3.	Perform Operations on the Sublist:
	•	Display the sliced sublist.
	•	Calculate and display the sum of the sublist elements.
	•	Calculate and display the average of the sublist elements.
	•	Find and display the maximum element in the sublist.
	4.	Allow Multiple Operations:
	•	After displaying the results, ask if the user wants to perform another slicing operation or exit.

Example :

Enter integers separated by space: 10 20 30 40 50 60
Enter start index: 2
Enter end index: 5

Sliced Sublist: [30 40 50]
Sum: 120
Average: 40.00
Maximum: 50

Do you want to slice again? (yes/no): yes
Enter start index: 0
Enter end index: 3

Sliced Sublist: [10 20 30]
Sum: 60
Average: 20.00
Maximum: 30

Do you want to slice again? (yes/no): no
Exiting the application. Goodbye!
