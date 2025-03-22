# Concurrent CSV Aggregator

We are reading a csv file and spiriting the file into n parts and
processing each part concurrently and then sending the processed data through channel and merging all the data

### Input
It takes a csv file as input
"transaction.csv" is the input here
the file should be in same directory as the exe file

### work1
File reader

Reads the file

### work2
Split the file to n parts

### works 3 
process the eacgh part cuncurrently 

### works 4
merge all the proceesed data 

## final out put

the list of all the username and amount
user_004    200
user_002    450.75
user_001    80
user_003    140.25


