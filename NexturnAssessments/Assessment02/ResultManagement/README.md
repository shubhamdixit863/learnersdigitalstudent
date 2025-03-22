ResultManagement
OBJECTIVE :
The application is used to garde and print the garding details of students from engineering and arts streams.
FEATURES:
1.Nested structures are used to enhance readability of code
2.Const variables are used
3.Interfaces are used to implement methods
EXPLANATION
1.models/student.go
 contains student struct and grade interface responsible for grading and print details
 2.models/engineering.go
 contains engineering student struct and implements calculategrade()  and getdetails() method
 3.models/arts.go
 contains arts student student struct and implements calculategrade()  and getdetails() method
 4.services/result.go
 contains processresults()to display results 
 5.cmd/main.go
 It initializes student data and calls the functions

 ---inorder to run application download the zip folder and run "go run cmd/main.go " in the terminal---
