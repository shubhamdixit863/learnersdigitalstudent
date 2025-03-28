package utilities

const (
	FilePath      = "../data"
	Slash         = "/"
	ErrorDirOpen  = "Error in scanning the directory"
	TakeInput     = "Please select the mode(1-4)"
	Modes         = "1. Line Filter Mode\n2. Word Count Mode\n3. API Call Mode\n4. Retryable API Call Mode"
	Keyword       = "Enter the word you want to search: "
	EmptyString   = ""
	ModeOne       = "The following are the lines containing the word: "
	ModeOneFail   = "No line contains the word: "
	ModeTwo       = "Total word count across all lines: "
	ModeThree     = "The following are the lines and their returned status code: "
	Colon         = ": "
	ModeThreeFail = "Error in Calling API: "
	ModeFourFail  = "maximum retries reached"
	URL           = "https://httpbin.org/post"
	ContentType   = "text/plain"
	DefaultError  = "Not a valid mode, please select a valid mode"
)
