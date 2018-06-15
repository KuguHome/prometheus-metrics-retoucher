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

Making it Runnable From the Command Line:
Compile the program with the following:
go build relabeler.go

This will make an executable, ‘relabeler’. After, the program has to be copied to the system path, which can be done by copying to /usr/local/bin:
cp stringparse /usr/local/bin


Example:
This is a line in a file called node.prom before and after being run through the script. The script can be called through the command line as follows:
cat path/node.prom.txt | relabeler --label 123=456 --label abc=def --label Austin=Li --label one=more > node-relabeled.prom.txt

Input:
go_gc_duration_seconds{quantile="0"} 7.091e-06

Output:
go_gc_duration_seconds{abc="def",Austin="Li",one="more",123="456",quantile="0"} 7.091e-06