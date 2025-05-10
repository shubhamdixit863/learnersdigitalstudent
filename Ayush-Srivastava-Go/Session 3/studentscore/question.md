# Problem Statement: Build a Student Score Management CLI Application in Golang (Without Structs)

Overview:

Create a Command-Line Interface (CLI) application in Golang that allows users to add, view, update, and delete student scores .

    1.	Add a Student:

    •	Prompt the user to enter a student name and score.

    •	Store names in a []string slice and scores in a []int slice.

    •	Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0]).

    2.	View All Students and Scores:

    •	Display all students with their scores.

    •	Handle cases when there are no students.

    3.	Update a Student’s Score:

    •	Ask for the student’s name and new score.

    •	Update the corresponding score in the scores slice.

    4.	Delete a Student:

    •	Prompt for the student’s name and remove both the name and corresponding score.

    5.	Exit the Application:

    •	Provide an option to terminate the program.

Features to Implement:

    •	Use two slices:

    •	names []string for storing student names.

    •	scores []int for storing corresponding scores.

    •	Create functions for:

    •	Adding a student (addStudent)

    •	Viewing all students (viewStudents)

    •	Updating scores (updateScore)

    •	Deleting students (deleteStudent)

    •	Use for loops to:

    •	Iterate through slices for display, update, and deletion operations.

    •	Use basic datatypes (string, int) for names and scores.

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
