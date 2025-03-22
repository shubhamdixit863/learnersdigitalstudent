package filewalker

import (
    "os"
    "path/filepath"
)


func WalkDirectory(dir string) ([]string, error) {
    var files []string

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return files, nil
}