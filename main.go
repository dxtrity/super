package main

import (
  "os"
  "fmt"
  "log"

  "dxtrity/super/help"
  "dxtrity/super/initialise"
)

func main() {
  args := os.Args[1:]
  
  if len(args) == 0 {
    help.OutputHelpCommand()
  } else if len(args) >= 1 {
    switch args[0] {
    case "init":
      err := initialise.InitEmptyRepo("./")
      if err != nil {
        log.Fatal(err)
      }
    case "commit":
      fmt.Println("This is not implemented yet!")
    case "add":
      fmt.Println("This is not implemented yet!")
    default:
      fmt.Println("This is not a valid command!")
    }
  }
}
