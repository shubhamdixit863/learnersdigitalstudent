package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"parallel_file_hasher/internal/models"
	"parallel_file_hasher/internal/services/filewalker"
	"parallel_file_hasher/internal/services/hasher"
	"parallel_file_hasher/internal/utils"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	logger, err := initLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, utils.ErrInitLogFmt, err)
		return utils.ExitCodeError
	}

	config, err := parseConfig()
	if err != nil {
		logger.Printf(utils.ErrProcessFmt, err)
		return utils.ExitCodeError
	}

	err = processFiles(config, logger)
	if err != nil {
		logger.Printf(utils.ErrProcessFmt, err)
		return utils.ExitCodeError
	}

	return utils.ExitCodeSuccess
}

func initLogger() (*log.Logger, error) {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	return logger, nil
}

func parseConfig() (*models.Config, error) {
	configPath := utils.GivenConfigPath

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf(utils.ErrConfigFileNotFound, configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrFailedToReadConfig, err)
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf(utils.ErrFailedToParseConfig, err)
	}

	if config.DirPath == "" {
		return nil, fmt.Errorf(utils.ErrDirPathRequired)
	}

	if config.Workers <= 0 {
		config.Workers = utils.DefaultWorkers
	}

	return &config, nil
}

func processFiles(config *models.Config, logger *log.Logger) error {
	walker := filewalker.NewFileWalker(config.DirPath)

	files, err := walker.Walk()
	if err != nil {
		return fmt.Errorf(utils.ErrReadDirFmt, err)
	}

	hasherService := hasher.NewHasherService(config.Workers)

	results, err := hasherService.ProcessFiles(files)
	if err != nil {
		return err
	}

	for _, result := range results {
		if result.Error != nil {
			logger.Printf("Error processing %s: %v", result.File, result.Error)
			continue
		}
		fmt.Printf("%s: %s\n", result.File, result.Hash)
	}

	return nil
}
