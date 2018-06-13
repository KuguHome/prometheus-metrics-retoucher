package main

import (
  "fmt"
  //"io"
  "bufio"
  //"os"
  "gopkg.in/alecthomas/kingpin.v2"
  )

  var (
    label = kingpin.Flag("label", "Add a custom label.").Strings()
  )

func main() {
  //parse -label from command line
  kingpin.Parse()
  //so now label is a "slice" (dynamic array kinda thing) containing
  //all the desired custom labels (called by doing --label labelname)

  //read file from stdin
  scanner := bufio.NewScanner(os.Stdin)
  var line String
  line := scanner.Scan()

  for line {
    if scanner.Text()[0] == "#" {
      fmt.Println("%s", scanner.Text())
      line := scanner.Scan()
    } else {
      
    }

  }

  //error catch block?
  if err := scanner.Err(); err != nil {
	   log.Println(err)
  }
}
