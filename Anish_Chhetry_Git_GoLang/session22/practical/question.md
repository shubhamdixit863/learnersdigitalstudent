// use net http package
// you have to create two  handlers ,which gives the file list
// and another gives the file data as per the name passed

// You have to create ,log middleware that logs following things
// first logs all things about request in structure format ,agent ,header
// is you have to create a map with the user list and the files they can access
// if the user tries to access any  other file other than what they have access too ,
//send 405 status forbidden