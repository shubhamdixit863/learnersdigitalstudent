package models

type LineWithMetadata struct {
	FileName string
	LineNum  int
	Line     string
}

type ProcessResult struct {
	FileName   string
	LineNum    int
	Line       string
	Keep       bool
	WordCount  int
	APISuccess bool
	APIStatus  string
	Error      error
}