package main

import (
  "strings"
  "os"

  "gopkg.in/alecthomas/kingpin.v2"

  dto "github.com/prometheus/client_model/go"
  
  "github.com/prometheus/common/expfmt"
  "github.com/golang/protobuf/proto"
  )

  //set up the --label flag
  var (
    flagArgs = kingpin.Flag("label", "Add a label and value in
      the form \"<label>=<value>\".").Strings()
  )

func main() {
  //parses command line flags
  kingpin.Parse()

  //creates TextParser and parses text into metrics
  var parser expfmt.TextParser
  parsedFamilies, _ := parser.TextToMetricFamilies(os.Stdin)

  //validPairs is a slice of POINTERS
  var validPairs []*dto.LabelPair

  //parses the flag arguments. Ignores if not the correct format
  for _, elem := range *flagArgs {
    eqInd := strings.Index(elem, "=")
      if eqInd != -1 {
        pair := &dto.LabelPair{
            Name:  proto.String(elem[:eqInd]),
            Value: proto.String(elem[eqInd+1:]),
          }
        validPairs = append(validPairs, pair)
      }
    }

  //appends the valid pairs to the metrics and writes to StdOut
  for _, metricFamily := range parsedFamilies {
    for _, metric := range metricFamily.Metric {
      metric.Label = append(metric.Label, validPairs...)
    }
    expfmt.MetricFamilyToText(os.Stdout, metricFamily)
  }
}
