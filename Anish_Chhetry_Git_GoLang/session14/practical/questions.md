
You need to implement a concurrent log processing system in Go that meets the following requirements:

1. Log Generation (Producers)

• Multiple goroutines should generate logs concurrently.

• Each log entry follows this format:

"[TIMESTAMP] LEVEL: MESSAGE"

• Example logs:

"[2025-03-10 12:00:00] INFO: User logged in"
"[2025-03-10 12:01:00] ERROR: Database connection failed"
"[2025-03-10 12:02:00] WARN: Disk usage is high"


2. Log Processing (Workers)

• Use a worker pool of goroutines to process logs.

• Extract information into a LogEntry struct with fields:

type LogEntry struct {
Timestamp time.Time
Level     string
Message   string
}


• Processed logs should be sent to another channel.



3. Aggregation & Storage

• Store processed log entries in a map, where keys are log levels (INFO, ERROR, WARN) and values are the count of occurrences.

• Ensure safe concurrent access to the map using synchronization mechanisms.

4. Concurrency Mechanisms to Use

• Goroutines for log generation and processing.

• Channels for communication between goroutines.

• WaitGroups (sync.WaitGroup) for worker synchronization.

• Mutex (sync.Mutex) to protect the shared map from race conditions.






Expected Output



At the end of execution, print the log level count summary:

Log Level Count Summary: map[INFO:8 ERROR:5 WARN:7]





This paste expires in <8 hours. Public IP access. Share whatever you see with others in seconds with  Context. Terms of ServiceReport this