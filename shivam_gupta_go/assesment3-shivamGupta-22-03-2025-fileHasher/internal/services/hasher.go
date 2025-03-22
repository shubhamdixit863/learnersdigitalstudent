package services

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"fileHasher/internal/model"
	"fileHasher/internal/utility"
)


func ComputeHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hasher := sha256.New()
    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }

    return hex.EncodeToString(hasher.Sum(nil)), nil
}


func HashFiles(filePaths []string) []model.HashResult {
    results := make([]model.HashResult, len(filePaths))
    resultChan := make(chan model.HashResult)

    
    for _, filePath := range filePaths {
        go func(path string) {
            hash, err := ComputeHash(path)
            if err != nil {
                hash = utility.ErrSimple
            }
            resultChan <- model.HashResult{File: path, Hash: hash}
        }(filePath)
    }

    
    for i := 0; i < len(filePaths); i++ {
        results[i] = <-resultChan
    }

    defer close(resultChan)
    return results
}