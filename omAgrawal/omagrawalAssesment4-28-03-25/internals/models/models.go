package models

type ProcessedData struct {
	FileName   string
	Data       string
	LineNumber int
}

func NewProcessedData(fileName string, data string, lineNumber int) *ProcessedData {
	return &ProcessedData{fileName, data, lineNumber}
}
