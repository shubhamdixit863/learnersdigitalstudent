# You are developing an Employee Management System where employee records are stored in JSON format. Your system needs to handle multiple employee records, process them, and generate summary reports.

Requirements:

1. Define Structs with Nested Structures:

• Create a struct Employee with the following fields:

• ID (int)

• Name (string)

• Age (int)

• Department (string)

• Salary (float64)

• Address (nested struct) with:

• Street (string)

• City (string)

• ZipCode (string)

• Skills (slice of strings)

• Use struct tags to ensure correct JSON marshalling/unmarshalling.

2. Process a JSON File Containing Multiple Employee Records:

• Read employee data from a JSON file (employees.json) which contains an array of employees.

• Unmarshal this JSON into a slice of Employee structs.

3. Implement Data Processing Logic:

• Increase Salary: Increase the salary of employees under 30 years old by 15% and for employees above 50 years by 10%.

• Sort Employees: Sort employees by Salary in descending order (highest salary first).

• Filter Employees by Department: Allow filtering employees by a specific department.

4. Generate a Summary Report:

• Marshal the updated employee list back to JSON (updated_employees.json).

• Print the total number of employees per department as a summary.

Input JSON (employees.json) Example:

[

    {

        "id": 101,

        "name": "Alice Johnson",

        "age": 28,

        "department": "Engineering",

        "salary": 75000.50,

        "address": {

            "street": "123 Main St",

            "city": "San Francisco",

            "zipCode": "94107"

        },

        "skills": ["Go", "Kubernetes", "Docker"]

    },

    {

        "id": 102,

        "name": "Robert Brown",

        "age": 55,

        "department": "Finance",

        "salary": 85000.75,

        "address": {

            "street": "456 Elm St",

            "city": "New York",

            "zipCode": "10001"

        },

        "skills": ["Accounting", "Excel", "Financial Analysis"]

    }

]

Expected Output JSON (updated_employees.json) Example (After Processing):

[

    {

        "id": 102,

        "name": "Robert Brown",

        "age": 55,

        "department": "Finance",

        "salary": 93500.83,

        "address": {

            "street": "456 Elm St",

            "city": "New York",

            "zipCode": "10001"

        },

        "skills": ["Accounting", "Excel", "Financial Analysis"]

    },

    {

        "id": 101,

        "name": "Alice Johnson",

        "age": 28,

        "department": "Engineering",

        "salary": 86250.58,

        "address": {

            "street": "123 Main St",

            "city": "San Francisco",

            "zipCode": "94107"

        },

        "skills": ["Go", "Kubernetes", "Docker"]

    }

]

Summary Report Example (Console Output):

Total Employees by Department:

Engineering: 1

Finance: 1
