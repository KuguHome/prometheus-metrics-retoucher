package main

import (
  "strings"
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  dto "github.com/prometheus/client_model/go"
  "github.com/prometheus/common/expfmt"
  "github.com/golang/protobuf/proto"
  )

  var (
    str = kingpin.Flag("label", "Add a custom label.").Strings()
  )

func main() {
  kingpin.Parse()

  var parser expfmt.TextParser
  parsedFamilies, _ := parser.TextToMetricFamilies(os.Stdin)

  var validPairs []*dto.LabelPair
  for _, elem := range *str {
    eqInd := strings.Index(elem, "=")
      if eqInd != -1 {
        pair := &dto.LabelPair{
            Name:  proto.String(elem[:eqInd]),
            Value: proto.String(elem[eqInd+1:]),
          }
        validPairs = append(validPairs, pair)
      }
    }

  for _, mf := range parsedFamilies {
    for _, m := range mf.Metric {
      m.Label = append(m.Label, validPairs...)
    }
  }

  for _, mf := range parsedFamilies {
    expfmt.MetricFamilyToText(os.Stdout, mf)
  }
}
