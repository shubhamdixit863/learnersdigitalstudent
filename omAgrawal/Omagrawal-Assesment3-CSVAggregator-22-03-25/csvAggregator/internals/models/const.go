package models

const (
	FileName = "transactions.csv"
	N        = 3
)
const (
	FileReadError      = "Error in Reading File"
	Fileprocessed      = "File is succesfully processed"
	FileReadSuccesfull = "File Read Successful"
	SplitSuccessFull   = "Split of data Successful"
	FinalOutput        = "Successfully transaction File is processed and the final putput is "
)

var SplitFileMapp = make(map[int][][]string)
