You need to design a Task Processing System in Golang that can execute various types of tasks while handling errors gracefully. The system should be built using structs, interfaces, and custom error handling.



# Requirements:

## 1. Task Definition using Interfaces and Structs:

• Define a Task interface with a method Run() error.

• Implement multiple types of tasks as structs (e.g., FileProcessingTask, DataValidationTask, APICallTask).

• Each task should provide its own implementation of Run().

## 2. Task Manager:

• A TaskManager struct should maintain a list of tasks.

• It should execute tasks sequentially (one after another).

• If a task fails, it should retry up to 3 times before marking it as failed.

## 3. Error Handling:

• Implement custom error types like TaskError and ValidationError.

• Use fmt.Errorf() with error wrapping for context.

• Use errors.Is() or errors.As() for checking error types.

## 4. Defer and Resource Cleanup:

• Ensure that resource cleanup (e.g., closing a file, network connection) is handled using defer.

## 5. Logging and Reporting:

• Implement structured logging for successful tasks and failed tasks.

• Store failed task details in a log file or a slice for later review.



## Bonus Features (Optional):

• Introduce task dependencies, where one task must complete before another starts.

• Add a configuration mechanism to allow retry limits to be adjustable.

• Save failed tasks to a JSON or CSV file for debugging.

