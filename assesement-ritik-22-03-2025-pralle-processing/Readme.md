<!-- Walk Directory  The filewalker package finds all files in the given directory.

Start Goroutines  The hasher package launches a Goroutine for each file to compute the hash.

Compute Hashes  Each Goroutine reads the file content and calculates the SHA-256 hash.

Send Results  The hash results are sent through a channel to the main function.

Collect & Display  The main function gathers the results and prints them.



Start
 |
 |--- WalkDirectory (Find files in the directory)
 |
 |--- Loop through files
       |       
       |--- Spawn Goroutine for each file
       |       |
       |       |--- Compute SHA-256 Hash
       |       |--- Send result via channel
       |
       |--- Collect results from channel
       |
       |--- Print results
       |
       --- End -->