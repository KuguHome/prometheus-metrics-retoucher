package main

import (
  "fmt"
  //"io"
  //"bufio"
  //"os"
  "gopkg.in/alecthomas/kingpin.v2"
  )

  var (
    label = kingpin.Flag("label", "Add a custom label.").Strings()
  )

func main() {
  //parse -label from command line
  kingpin.Parse()

  //read file from stdin
  //scanner := bufio.NewScanner(os.Stdin)
}
