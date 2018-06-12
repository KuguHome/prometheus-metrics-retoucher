## Prometheus relabeler

### General objective
Implement a program which ingests raw prometheus scrape data, adds some custom labels and outputs it again.

### Detailed requirements (first step)
- An example call to the program should look like this (name and syntax of the arguments is only an example, use what fits best to the client libraries):
`cat node.prom | relabeler -label instance=some_instance -label job=some_job  > node-relabeled.prom`
- The program should accept data from STDIN, so it can be used on Linux via pipe (|) 
- The program should output the data to STDOUT
- An example input/output pair should look like this:
Input:
```
go_goroutines 9
```

Output:
```
go_goroutines{instance="some_instance",job="some_job"} 9
```

### Additional requirements
These can be implemented for later versions (nice to have)
- add argument to filter out/drop certain metrics (`relabeler -filter http_requests_total,http_request_duration_microseconds`)
- add argument to read in file (`relabeler -in node.prom`)
- add argument to output to file (`relabeler -out node-relabeled.prom`)
- add argument to read in directory (`relabeler -dir scrapes/`)



### Development environment in Go
As an editor you could use https://atom.io/
Follow the instructions at https://golang.org/doc/install
Video about it: https://www.youtube.com/watch?v=sNogq_98wV0
Additional tools for Atom/Go: https://rominirani.com/setup-go-development-environment-with-atom-editor-a87a12366fcf

### Documentation
* Short introduction to prometheus in general https://www.youtube.com/watch?v=WUkNnY65htQ
* Longer talk about some Prometheus details (in conjunction with Docker) https://www.youtube.com/watch?v=PDxcEzu62jk
* Library for implementation in Go https://github.com/prometheus/client_golang
* Library for implementation in Python, including some examples  https://github.com/prometheus/client_python
* Prometheus documentation with information about concepts etc. https://prometheus.io/docs/concepts/data_model/
