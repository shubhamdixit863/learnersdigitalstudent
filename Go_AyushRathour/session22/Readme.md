using net http package
I have created two  handlers ,
1. which gives the file list
2. another gives the fileData as per the name passed

I have created a log middleware that logs following things
1. first logs all things about request in structure format 
2. agent 
3. header

 I have created a map with the user list and the files they can access
* if the user tries to access any  other file other than what they have access to, send 405 status forbidden