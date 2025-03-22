-------------------------------------- Concurrent Log Analyzer --------------------------------------

* Overview
This project is a tool that reads a large `.log` file and computes how many times each log level 
(e.g., `INFO`, `ERROR`, `DEBUG`, etc.) appears. The log file is processed in chunks (e.g., 1000 lines per chunk), 
and each chunk is handled concurrently using goroutines.

* Features
- Reads large `.log` files efficiently.
- Splits log files into chunks and processes each chunk concurrently.
- Counts occurrences of different log levels (`INFO`, `ERROR`, `DEBUG`, etc.).
- Uses buffered channels to collect and aggregate results efficiently.

* Project Structure
concurrent-log-analyzer/
│── main.go           # Entry point of the application
│── reader.go         # Reads the log file and splits it into chunks
│── analyzer.go       # Processes each chunk and counts log levels
│── aggregator.go     # Collects results and prints the final summary
│── README.md         # Project documentation

* Creating executable
  go build main.go

* Run the Program
  .\main.exe

* Example Output

Log Level Counts:
WARNING :  3
DEBUG :  6
ERROR :  6
INFO :  9

