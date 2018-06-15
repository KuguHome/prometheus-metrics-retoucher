Setup:
This script uses the language Golang. Go to the following website for installation instructions:
https://golang.org/doc/install

You will also need to perform:
go get github.com/prometheus/client_model/go

Script Details:
The program reads from STDIN. It then parses the text into metrics, adds labels as requested through the command line, puts back together formatted text with the new labels, and writes it to STDOUT.

Command Line:
--label <label>=<value> 
	The label-value pair <label>=<value> is added to the incoming text in the correct 	format. --label can be called an arbitrary number of times.

Example:
This is a line in a .prom file before and after being run through the script:

Input:
go_gc_duration_seconds{quantile="0"} 7.091e-06

Output: