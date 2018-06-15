package main

import (
  "os"

  "gopkg.in/alecthomas/kingpin.v2"

  dto "github.com/prometheus/client_model/go"

  "github.com/prometheus/common/expfmt"
  "github.com/golang/protobuf/proto"
  )

  //set up the --label flag
  var (
    flagArgs = kingpin.Flag("label", "Add a label and value in" +
      "the form \"<label>=<value>\".").StringMap()
  )

func main() {
  //parses command line flags into a key=value map
  kingpin.Parse()

  //creates TextParser and parses text into metrics
  var parser expfmt.TextParser
  parsedFamilies, _ := parser.TextToMetricFamilies(os.Stdin)

  //validPairs is a slice of POINTERS
  var validPairs []*dto.LabelPair

  //converts map into LabelPair slice
  for key, value := range *flagArgs {
        validPairs = append(validPairs, &dto.LabelPair{
					Name:  proto.String(key),
					Value: proto.String(value),
				})
      }

  //appends the valid pairs to the metrics and writes to StdOut
  for _, metricFamily := range parsedFamilies {
    for _, metric := range metricFamily.Metric {
      metric.Label = append(metric.Label, validPairs...)
    }
    expfmt.MetricFamilyToText(os.Stdout, metricFamily)
  }
}
