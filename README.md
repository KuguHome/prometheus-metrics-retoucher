## Prometheus relabeler/retoucher

### Details
This program is made for the purpose of retouching metrics from exporters before ingesting them into [Prometheus](https://prometheus.io).

The program reads from STDIN or a file. It then parses the text into metrics, adds labels or drops metrics as requested through the command line, puts back together formatted text with the new labels, and writes it to STDOUT or a file.

### Command Line
`-a, --add-label <label>=<value>`
The label-value pair \<label\>=\<value\> is added to the incoming text in the correct format. Can be called an arbitrary number of times.

`-d, --drop-metric some_metric`
The metric given by some_metric is dropped. Can be called an arbitrary number of times.

`--in file_name`
Read in from file "file_name"

`--out file_name`
Write out to a file "file_name"

### Example
This is a line in a file called node.prom before and after being run through the script. The script can be called through the command line as follows:
```
cat node.prom.txt | relabeler --add-label instance=some_instance -a job=some_job  > node-relabeled.prom.txt
```
or
```
relabeler --in node.prom.txt --out node-relabeled.prom.txt --add-label instance=some_instance -a job=some_job -d node_network_transmit_multicast
```

Input (node.prom.txt):
```
go_gc_duration_seconds{quantile="0"} 7.091e-06
```

Output (node-relabeled.prom.txt):
```
go_gc_duration_seconds{instance="some_instance",job="some_job",quantile="0"} 7.091e-06
```

### Applications
The tool can be used to retouch scrapes from different distributed node_exporter instances, before pushing them to the [Pushgateway](https://github.com/prometheus/pushgateway). Normally, there would be a collision between the scrape's metrics and pushgateway's internal metrics, so before pushing, you can drop the colliding metrics via this tool:
```
cat node.prom.txt | ./relabeler \
--drop-metric go_memstats_last_gc_time_seconds \
--drop-metric go_goroutines --drop-metric go_memstats_other_sys_bytes \
--drop-metric go_gc_duration_seconds \
--drop-metric process_virtual_memory_bytes \
--drop-metric go_memstats_heap_inuse_bytes \
--drop-metric process_open_fds \
--drop-metric go_memstats_heap_alloc_bytes \
--drop-metric go_threads \
--drop-metric go_memstats_mcache_inuse_bytes \
--drop-metric process_max_fds \
--drop-metric go_memstats_alloc_bytes \
--drop-metric http_response_size_bytes \
--drop-metric process_start_time_seconds \
--drop-metric go_memstats_heap_released_bytes \
--drop-metric go_memstats_sys_bytes \
--drop-metric go_memstats_heap_idle_bytes  \
--drop-metric process_resident_memory_bytes  \
--drop-metric go_memstats_mcache_sys_bytes  \
--drop-metric go_memstats_frees_total  \
--drop-metric go_memstats_heap_objects \
--drop-metric go_memstats_next_gc_bytes  \
--drop-metric go_memstats_buck_hash_sys_bytes \
--drop-metric go_memstats_stack_sys_bytes \
--drop-metric go_memstats_heap_sys_bytes \
--drop-metric go_memstats_mspan_inuse_bytes \
--drop-metric go_memstats_gc_cpu_fraction \
--drop-metric go_memstats_stack_inuse_bytes \
--drop-metric http_request_duration_microseconds \
--drop-metric go_memstats_mspan_sys_bytes \
--drop-metric go_info \
--drop-metric go_memstats_gc_sys_bytes \
--drop-metric http_requests_total \
--drop-metric go_memstats_lookups_total \
--drop-metric process_cpu_seconds_total \
--drop-metric go_memstats_mallocs_total \
--drop-metric go_memstats_alloc_bytes_total \
--drop-metric http_request_size_bytes \
| curl --data-binary @- http://localhost:9091/metrics/job/node/instance/remote-machine
```

Or alternatively instead of using the Pushgateway to merge different distributed exporter scrapes, you could use [node_exporter's textfile collector](https://github.com/prometheus/node_exporter/blob/master/README.md#textfile-collector). Use the `--add-label` argument of this tool to add a different label to the metrics in each file to be able to discern them later in prometheus. But this will NOT work with scrapes of other node_exporters, since the metrics from the .prom files will collide with node_exporter's own metrics. A way to go may be to deactivate ALL collectors EXCEPT the textfile collector.

### Development/Build Setup
This program uses the language Golang. Go to the following website for installation instructions:
```
https://golang.org/doc/install
```

You will also need to perform:
```
go get github.com/prometheus/client_model/go
```

### Making it Runnable From the Command Line
Compile the program with the following:
```
go build relabeler.go
```

This will make an executable, ‘relabeler’. After, the program can be copied to the system path, which can be done by copying to /usr/local/bin:
```
cp relabeler /usr/local/bin
```

