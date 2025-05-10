package model

type HashResult struct {
	File string
	Hash string
}

func NewHashRsult(file string, hash string) *HashResult {
	return &HashResult{
		File: file,
		Hash: hash,
	}
}
