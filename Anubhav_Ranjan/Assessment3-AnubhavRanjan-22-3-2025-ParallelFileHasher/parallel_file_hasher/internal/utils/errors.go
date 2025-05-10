package utils

const (
	ErrInvalidDirPath = "provided directory path is not a directory"
	ErrReadDir        = "error reading directory: %w"
	ErrOpenFileFmt    = "failed to open file: %w"
	ErrReadFileFmt    = "failed to read file: %w"
	ErrComputeHash    = "failed to compute hash: %w"
)

const (
	DefaultWorkers  = 5
	ErrInitLogFmt   = "failed to initialize logger: %v"
	ErrReadDirFmt   = "failed to read directory: %v"
	ErrProcessFmt   = "failed to process files: %v"
	ExitCodeSuccess = 0
	ExitCodeError   = 1
)

const (
	GivenConfigPath        = "../config/config.json"
	ErrConfigFileNotFound  = "config file not found: %s"
	ErrFailedToReadConfig  = "failed to read config file: %v"
	ErrFailedToParseConfig = "failed to parse config file: %v"
	ErrDirPathRequired     = "dirPath is required in config file"
)
