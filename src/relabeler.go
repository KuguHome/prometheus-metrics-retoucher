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
    inFileFlagArg = kingpin.Flag("in", "Accepts a file").File();
    outFileFlagArg = kingpin.Flag("out", "Writes to a file").String(); //string because has to create the file
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

  //delete metrics requested to be dropped
  for _, name := range *dropFlagArgs {
    delete(parsedFamilies, name)
  }

  //appends the valid pairs to the metrics and write everything to STDOUT
  if *outFileFlagArg == "" {
    for _, metricFamily := range parsedFamilies {
      for _, metric := range metricFamily.Metric {
        metric.Label = append(metric.Label, validPairs...)
      }
      expfmt.MetricFamilyToText(os.Stdout, metricFamily)
    }
  } else {
      file, _ := os.Create(*outFileFlagArg)
      for _, metricFamily := range parsedFamilies {
        for _, metric := range metricFamily.Metric {
          metric.Label = append(metric.Label, validPairs...)
        }
        expfmt.MetricFamilyToText(file, metricFamily)
      }
  }

  //this is here as a reminder that writing to STDOUT might need its own
  //loop in case there's something else to do after appending valid pairs
  //for _, metricFamily := range parsedFamilies {
  //  expfmt.MetricFamilyToText(os.Stdout, metricFamily)
  //}
}
