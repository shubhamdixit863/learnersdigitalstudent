package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"concurrent_file_processor/internal/config"
	"concurrent_file_processor/internal/errors"
	"concurrent_file_processor/internal/services/fileprocessor"
)

const (
	defaultDirectory = "testfiles"
	defaultMode      = "filter"
	defaultKeyword   = "error"
)

func main() {
	directory := flag.String("dir", defaultDirectory, "Directory to process")
	mode := flag.String("mode", defaultMode, "Processing mode (filter/wordcount/apicall/retryapi)")
	keyword := flag.String("keyword", defaultKeyword, "Keyword for line filtering")
	flag.Parse()

	cfg := config.NewConfig()
	cfg.Directory = *directory
	cfg.Mode = config.ProcessingMode(*mode)
	cfg.Keyword = *keyword

	if err := runProcessor(cfg); err != nil {
		log.Fatal(errors.LogError(err))
	}
}

func runProcessor(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	processor, err := fileprocessor.NewFileProcessor(cfg)
	if err != nil {
		return fmt.Errorf(errors.ErrProcessingFailed, err.Error())
	}

	results, err := processor.Process(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Processing Results:")
	for _, result := range results {
		fmt.Println(result)
	}

	return nil
}
