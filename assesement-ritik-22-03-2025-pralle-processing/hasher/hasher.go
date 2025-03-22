package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"parallel_file_hasher/results"
)

func ComputeHash(filePath string, resultChan chan results.HashResult) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", filePath)
		resultChan <- results.HashResult{File: filePath, Hash: "ERROR"}
		return
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("Error reading file: %s\n", filePath)
		resultChan <- results.HashResult{File: filePath, Hash: "ERROR"}
		return
	}

	hashValue := hex.EncodeToString(hash.Sum(nil))
	resultChan <- results.HashResult{File: filePath, Hash: hashValue}
}
