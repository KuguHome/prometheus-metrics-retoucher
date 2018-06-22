package main

import (
  "os"
  "io"
  "bufio"

  "gopkg.in/alecthomas/kingpin.v2"

  dto "github.com/prometheus/client_model/go"

  "github.com/prometheus/common/expfmt"
  "github.com/golang/protobuf/proto"
  )

  //set up the flags
  var (
    labelFlagArgs = kingpin.Flag("add-label", "Add a label and value in the form \"<label>=<value>\".").PlaceHolder("<label>=<value>").Short('a').StringMap()
    dropFlagArgs = kingpin.Flag("drop-metric", "Drop a metric").PlaceHolder("some_metric").Short('d').Strings()
    inFileFlagArg = kingpin.Flag("in", "Read in a file").PlaceHolder("file_name").File();
    outFileFlagArg = kingpin.Flag("out", "Write to a File").PlaceHolder("file_name").String(); //string because has to create the file
    defaultDropFlag = kingpin.Flag("drop-default-metrics", "Drop default metrics").Bool();
  )

func main() {
  //parses command line flags into a key=value map
  kingpin.Parse()

  //creates TextParser and parses text into metrics
  var parser expfmt.TextParser
  var reader io.Reader
  if *inFileFlagArg == nil {
    reader = os.Stdin
  } else {
    reader = bufio.NewReader(*inFileFlagArg)
  }
  parsedFamilies, _ := parser.TextToMetricFamilies(reader)

  //validPairs is a slice of POINTERS
  var validPairs []*dto.LabelPair

  //converts map into LabelPair slice
  for key, value := range *labelFlagArgs {
        validPairs = append(validPairs, &dto.LabelPair{
					Name:  proto.String(key),
					Value: proto.String(value),
				})
      }

  //add the default drop metrics to the list of metrics to be dropped
  if *defaultDropFlag {
    *dropFlagArgs = append(*dropFlagArgs, []string{
      "go_memstats_last_gc_time_seconds",
      "go_goroutines",
      "go_memstats_other_sys_bytes",
      "go_gc_duration_seconds",
      "process_virtual_memory_bytes",
      "go_memstats_heap_inuse_bytes",
      "process_open_fds",
      "go_memstats_heap_alloc_bytes",
      "go_threads",
      "go_memstats_mcache_inuse_bytes",
      "process_max_fds",
      "go_memstats_alloc_bytes",
      "http_response_size_bytes",
      "process_start_time_seconds",
      "go_memstats_heap_released_bytes",
      "go_memstats_sys_bytes",
      "go_memstats_heap_idle_bytes",
      "process_resident_memory_bytes",
      "go_memstats_mcache_sys_bytes",
      "go_memstats_frees_total",
      "go_memstats_heap_objects",
      "go_memstats_next_gc_bytes",
      "go_memstats_buck_hash_sys_bytes",
      "go_memstats_stack_sys_bytes",
      "go_memstats_heap_sys_bytes",
      "go_memstats_mspan_inuse_bytes",
      "go_memstats_gc_cpu_fraction",
      "go_memstats_stack_inuse_bytes",
      "http_request_duration_microseconds",
      "go_memstats_mspan_sys_bytes",
      "go_info",
      "go_memstats_gc_sys_bytes",
      "http_requests_total",
      "go_memstats_lookups_total",
      "process_cpu_seconds_total",
      "go_memstats_mallocs_total",
      "go_memstats_alloc_bytes_total",
      "http_request_size_bytes"}...)
  }

  //delete metrics requested to be dropped
  for _, name := range *dropFlagArgs {
    delete(parsedFamilies, name)
  }

  //appends the valid pairs to the metrics and write out
  if *outFileFlagArg == "" {
      writeOut(parsedFamilies, validPairs, os.Stdout)
  } else {
      outFile, _ := os.Create(*outFileFlagArg)
      writeOut(parsedFamilies, validPairs, outFile)
  }
}

//rebuild the text with the new labels and write to writeTo
func writeOut(families map[string]*dto.MetricFamily, labelPairs []*dto.LabelPair, writeTo io.Writer) {
  for _, metricFamily := range families {
    for _, metric := range metricFamily.Metric {
      metric.Label = append(metric.Label, labelPairs...)
    }
    expfmt.MetricFamilyToText(writeTo, metricFamily)
  }
}
