package models

type HashResult struct {
	File  string
	Hash  string
	Error error
}

type Config struct {
	DirPath string
	Workers int
}
