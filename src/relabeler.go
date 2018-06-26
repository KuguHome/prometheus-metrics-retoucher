package main

import (
  "os"
  "io"
  "bufio"
  "io/ioutil"
  "log"
  "strings"

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
    inDirFlagArg = kingpin.Flag("in-dir", "Read in a directory").PlaceHolder("dir_name").String();
    defaultFlags = []string{
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
      "http_request_size_bytes"}
  )

func main() {
  //parses command line flags into a key=value map
  kingpin.Parse()

  var writer io.Writer
  if *outFileFlagArg != "" {
    var err error
    writer, err = os.Create(*outFileFlagArg)
    if err != nil {
        log.Fatal(err)
    }
  } else {
    writer = os.Stdout
  }

  if (*inFileFlagArg == nil) && (*inDirFlagArg == "") {
    parseAndRebuild(os.Stdin, os.Stdout)
  } else {
    if *inFileFlagArg != nil {
      reader := bufio.NewReader(*inFileFlagArg)
      parseAndRebuild(reader, writer)
    }
    if *inDirFlagArg != "" {
      filesInfo, err := ioutil.ReadDir(*inDirFlagArg)
      if err != nil {
          log.Fatal(err)
      }
      for _, info := range filesInfo {
        //not sure if this if statement is necessary
        if strings.HasSuffix(info.Name(), ".prom") {
          reader, err := os.Open("" + *inDirFlagArg + "/" + info.Name())
          if err != nil {
              log.Fatal(err)
          }
          parseAndRebuild(reader, writer)
        }
      }
    }
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

//converts key-value map into LabelPair slice
func pairToSlice(pairs []*dto.LabelPair) []*dto.LabelPair {
  for key, value := range *labelFlagArgs {
        pairs = append(pairs, &dto.LabelPair{
					Name:  proto.String(key),
					Value: proto.String(value),
				})
      }
      return pairs;
}

func parseAndRebuild(readFrom io.Reader, writeTo io.Writer) {
    //creates TextParser and parses text into metrics
  var parser expfmt.TextParser

  parsedFamilies, _ := parser.TextToMetricFamilies(readFrom)

  //validPairs is a slice of POINTERS
  var validPairs []*dto.LabelPair

  //converts map into LabelPair slice
  validPairs = pairToSlice(validPairs)

  //add the default drop metrics to the list of metrics to be dropped
  if *defaultDropFlag {
    *dropFlagArgs = append(*dropFlagArgs, defaultFlags...)
  }

  //delete metrics requested to be dropped
  for _, name := range *dropFlagArgs {
    delete(parsedFamilies, name)
  }
  writeOut(parsedFamilies, validPairs, writeTo)
}
