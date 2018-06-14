package main

import (
  //"fmt"
  "strings"
  //"log"
  //"io"
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
  //parse -label from command line
  kingpin.Parse()
  //so now label is a "slice" (dynamic array kinda thing) containing
  //all the desired custom labels (called by doing --label labelname)

  //read file from stdin
  //scanner := bufio.NewScanner(os.Stdin)
  //var line bool
  //line = scanner.Scan()

  var parser expfmt.TextParser
  parsedFamilies, _ := parser.TextToMetricFamilies(os.Stdin)

  for _, elem := range *str {
    eqInd := strings.Index(elem, "=")
      if eqInd != -1 {
        for _, mf := range parsedFamilies {
          for _, m := range mf.Metric {
            m.Label = append(m.Label, &dto.LabelPair{
                Name:  proto.String(elem[:eqInd]),
                Value: proto.String(elem[eqInd+1:]),
            })
          }
        }
      }
    }
  for _, mf := range parsedFamilies {
    expfmt.MetricFamilyToText(os.Stdout, mf)
  }
}
