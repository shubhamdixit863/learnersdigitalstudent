package utils

const (
	FileProcessingSignal = "Processing files in directory:"
	FileProcessingError  = "Error processing files:"
	ModeProcessingError  = "Error processing mode:"
	SearchedFileFormat   = ".txt"
	FileReadingError     = "Error reading file:"
	FilteredText         = "Error"
	InvalidModeError     = "Invalid mode"
	TotalWords           = "Total words:"
	APIStatus            = "API Call status"
	Filtered             = "Filtered"
	url                  = "https://httpbin.org/post"
	ContentType          = "application/json"
	APISuccess           = "Success"
	APIFailed            = "Failed"
	APIRetying           = "Retrying API call"
	SearchedDirectory    = "./data"
)

var Modes = []string{"word_count", "filter", "api_call"}