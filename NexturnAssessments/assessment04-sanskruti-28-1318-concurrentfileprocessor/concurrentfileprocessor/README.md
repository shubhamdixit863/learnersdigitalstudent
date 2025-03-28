CONCURRENT FILE PROCESSOR
### objectives
The code reads multiple txt files and and processes them concurrently using FAN-OUT and FAN-IN concurrency pattern.

### Features
1.constant have been used wherever necessary to enhance the readability of the code
2.waitgroups ,mutex have been used to increased synchronicity of the code
3.Error handling is done incase error is encountered while opening ,reading files 

### aFOLDER STRUCTURE
1.cmd/main.go
Flags have been used to take input .Channel is initialized.
2.internal/reader/filereader.go
opens and reads the file .Converts the text to string byte slice.
3.internal/services/processor.go
Contains all the goroutines required for carrying out operations.


