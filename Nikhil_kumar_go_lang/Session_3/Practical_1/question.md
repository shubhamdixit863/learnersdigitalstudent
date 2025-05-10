Problem Statement: Build a Student Score Management CLI Application in Golang (Without Structs)

Overview:

Create a Command-Line Interface (CLI) application in Golang that allows users to add, view, update, and delete student scores .

1. Add a Student:
   • Prompt the user to enter a student name and score.
   • Store names in a []string slice and scores in a []int slice.
   • Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0]).
2. View All Students and Scores:
   • Display all students with their scores.
   • Handle cases when there are no students.
3. Update a Student’s Score:
   • Ask for the student’s name and new score.
   • Update the corresponding score in the scores slice.
4. Delete a Student:
   • Prompt for the student’s name and remove both the name and corresponding score.
5. Exit the Application:
   • Provide an option to terminate the program.

Features to Implement:
• Use two slices:
• names []string for storing student names.
• scores []int for storing corresponding scores.
• Create functions for:
• Adding a student (addStudent)
• Viewing all students (viewStudents)
• Updating scores (updateScore)
• Deleting students (deleteStudent)
• Use for loops to:
• Iterate through slices for display, update, and deletion operations.
• Use basic datatypes (string, int) for names and scores.

Student Score Manager:

1. Add Student
2. View Students
3. Update Student Score
4. Delete Student
5. Exit

Enter choice: 1
Enter Student Name: Alice
Enter Score: 85
Student added successfully!

Enter choice: 2
Students and Scores:

1. Alice - 85

Enter choice: 3
Enter Student Name to Update: Alice
Enter New Score: 90
Score updated successfully!

Enter choice: 4
Enter Student Name to Delete: Alice
Student deleted successfully!

Enter choice: 2
No students found.

Problem Statement: Build a Student Score Management CLI Application in Golang (Without Structs)

Overview:

Create a Command-Line Interface (CLI) application in Golang that allows users to add, view, update, and delete student scores .

1. Add a Student:
   • Prompt the user to enter a student name and score.
   • Store names in a []string slice and scores in a []int slice.
   • Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0]).
2. View All Students and Scores:
   • Display all students with their scores.
   • Handle cases when there are no students.
3. Update a Student’s Score:
   • Ask for the student’s name and new score.
   • Update the corresponding score in the scores slice.
4. Delete a Student:
   • Prompt for the student’s name and remove both the name and corresponding score.
5. Exit the Application:
   • Provide an option to terminate the program.

Features to Implement:
• Use two slices:
• names []string for storing student names.
• scores []int for storing corresponding scores.
• Create functions for:
• Adding a student (addStudent)
• Viewing all students (viewStudents)
• Updating scores (updateScore)
• Deleting students (deleteStudent)
• Use for loops to:
• Iterate through slices for display, update, and deletion operations.
• Use basic datatypes (string, int) for names and scores.

Student Score Manager:

1. Add Student
2. View Students
3. Update Student Score
4. Delete Student
5. Exit

Enter choice: 1
Enter Student Name: Alice
Enter Score: 85
Student added successfully!

Enter choice: 2
Students and Scores:

1. Alice - 85

Enter choice: 3
Enter Student Name to Update: Alice
Enter New Score: 90
Score updated successfully!

Enter choice: 4
Enter Student Name to Delete: Alice
Student deleted successfully!

Enter choice: 2
No students found.

Overview:

Create a CLI application in Golang that performs various operations on a list of integers provided by the user. The application should use the slicing operator (:) to create sublists and calculate statistics like sum, average, and maximum of the sliced portion.

Requirements:

1. Input List of Integers:
   • Prompt the user to enter a list of integers separated by spaces (e.g., 10 20 30 40 50).
   • Store these integers in a slice.
2. Slice the List:
   • Ask the user to input start and end indices.
   • Use the slice operator to create a sublist from the original list.
   • Validate that indices are within bounds and that start <= end.
3. Perform Operations on the Sublist:
   • Display the sliced sublist.
   • Calculate and display the sum of the sublist elements.
   • Calculate and display the average of the sublist elements.
   • Find and display the maximum element in the sublist.
4. Allow Multiple Operations:
   • After displaying the results, ask if the user wants to perform another slicing operation or exit.

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
