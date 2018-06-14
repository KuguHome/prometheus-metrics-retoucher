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
      the form <name>=<value>.").Strings()
  )

func main() {
  kingpin.Parse()

  var parser expfmt.TextParser
  parsedFamilies, _ := parser.TextToMetricFamilies(os.Stdin)

  var validPairs []*dto.LabelPair
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

  for _, metricFamily := range parsedFamilies {
    for _, metric := range metricFamily.Metric {
      metric.Label = append(metric.Label, validPairs...)
    }
    expfmt.MetricFamilyToText(os.Stdout, metricFamily)
  }
}
