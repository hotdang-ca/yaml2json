package main

import (
  "os"
  "fmt"
  "flag"
  "io/ioutil"
  "github.com/ghodss/yaml"
)

func main() {
  reversePtr := flag.Bool("r", false, "Convert from json to yaml")

  flag.Parse()

  fi, err := os.Stdin.Stat()
  if err != nil {
    panic(err)
  }
  if fi.Mode() & os.ModeNamedPipe == 0 {
    fmt.Fprintf(os.Stderr, "No stdin provided. Try again.\n")
    os.Exit(1)
  } else {
    var d []byte
    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
      panic(err)
    }
    if *reversePtr {
      d, err = yaml.JSONToYAML(bytes)
    } else {
      d, err = yaml.YAMLToJSON(bytes)
    }
    if err != nil {
      panic(err)
    }
    fmt.Printf("%s", string(d))
  }
}
