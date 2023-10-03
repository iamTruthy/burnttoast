This is another hands-on exercise that shows some go commands. This exercise runs through

1. creating variables with the following scopes:

○ package level
■ create outside of func main
■ use the
● var keyword
● const keyword

○ block level
■ inside func main
■ use the short declaration operator



2. using the terminal to make a go workspace
○  mkdir <name>
○  cd <name>
○  go mod init <somename>

●  Write a Hello World program
○ nano main.go
○ write go code
○ run your program
○ go run main.go



3. Building the code in our program to run on different OS's:
● build our program for windows
○ GOOS=windows go build
● build our program for mac
○ GOOS=darwin go build
● build our program for linux
○ GOOS=linux go build

(this could program include an executable go file for MacOS, Windows and Linux
 if the comands on no.3 are ran correctly)