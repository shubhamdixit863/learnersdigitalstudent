package utils

var (
	dirFiles       = "./internal/files/"
	dirData        = "./internal/data/"
	Files          = []string{dirFiles + "file1.txt", dirFiles + "file2.txt", dirFiles + "file3.txt", dirFiles + "file4.txt", dirFiles + "book.txt", dirData + "file1.txt", dirData + "file2.txt", dirData + "file3.txt"}
	FileProcessErr = "Error processing file %s: %v"
	FileOpenErr    = "failed to open file: %w"
	FileReadErr    = "error reading file: %w"
)