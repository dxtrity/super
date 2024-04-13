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
      pathExists, err := exists("./.super")
      if err != nil {
        log.Fatalf("[UNEXPECTED]: %v", err)
      }
      if !pathExists {
        err := initialise.InitEmptyRepo("./")
        if err != nil {
          log.Fatalf("[UNEXPECTED]: %v", err)
        }
      } else {
        fmt.Println("[EXPECTED]: Repository already exists.")
      }
    case "commit":
      fmt.Println("[EXPECTED]: This is not implemented yet!")
    case "add":
      fmt.Println("[EXPECTED]: This is not implemented yet!")
    default:
      fmt.Print("[EXPECTED]: Improper command usage:\n\tsuper <command> -<option> <value>")
    }
  }
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}
