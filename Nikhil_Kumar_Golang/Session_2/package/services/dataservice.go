package services


// if a function or a variable is in the lower case its not avaiable in the other package
func getData()  string{
	return "hello"
}
func GetData()  string{
	return "hello from service"
}