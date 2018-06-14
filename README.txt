Setup:
This script uses the language Golang. Go to the following website for installation instructions:
https://golang.org/doc/install

To install packages and dependencies, do:
go get

Script Details:
The program reads from STDIN. It then parses the text into metrics, adds labels as requested through the command line, puts back together formatted text with the new labels, and writes it to STDOUT.

Command Line:
--label <label>=<value> 
	The label-value pair <label>=<value> is added to the incoming text in the correct 	format. --label can be called an arbitrary number of times.