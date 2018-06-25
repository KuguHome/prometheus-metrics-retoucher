package main

import (
  "fmt"
  "strings"
  "log"
  //"io"
  "bufio"
  "os"
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
  var line bool
  line = scanner.Scan()

  for line {
    if strings.Index(scanner.Text(), "#") == 0 {
      fmt.Println(scanner.Text())
      line = scanner.Scan()
    } else {
      //if there are already labels, add custom label and commas
      braceInd := strings.Index(scanner.Text(), "{")
      if braceInd != -1 {
        fmt.Print(scanner.Text()[0:braceInd])
        for _, element := range label {
          fmt.Print(element.)
        }
      }
      fmt.Println(scanner.Text())
      line = scanner.Scan()
    }

  }

  //error catch block?
  if err := scanner.Err(); err != nil {
	   log.Println(err)
  }
}
