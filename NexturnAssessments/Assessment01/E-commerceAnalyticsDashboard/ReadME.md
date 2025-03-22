E-commerceAnalyticsDashboard

The code initially reads the transaction data that is stored in csv file.
The results have been clubbed together in struct format
Maps have been used to to store the key value pairs of product -revenue,product-quantity sold,customer-total amt spend
Error handling has been used incase csv file failes to open .
---FLOW OF THE PROGRAM---
1.main function reads the file path
2.It calls the processtransactions function present in internal package 
3.Processtransaction funcn --
3.1 initially it reads data from the csv file provided
3.2 the read data is stored in the form of 2D slice of records
3.3 maps are created 
3.4 each record is fetched from 2 D slice of records and split into individual componenets
3.5 map value is incremented
3.6 logic for desired output is applied
3.7 results are returned in form of struct
4.Error handling is done incase we encounter error while fetching results
5.Results are printed 
6.Inorder to run the code download the zip folder and run "go run cmd/main.go" in the terminal