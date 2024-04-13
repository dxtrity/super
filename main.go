/*
  QUICK NOTES ABOUT PROJECT:

  ALL ERRORS SHOULD HAVE "[EXPECTED]" OR "[UNEXPECTED] EMBEDDED"
  EASIER TO DEBUG BC WE KNOW IF IT'S US FUCKING UP OR AN EXTERNAL LIB FUCKING UP.#

  [EXPECTED] = purposefully thrown errors like not implemented messages.
  [UNEXPECTED] = actual errors like os can't read file etc.
*/

package main

import (
  "os"
  "fmt"
  "log"

  "dxtrity/super/help"
  "dxtrity/super/initialise"
)

func main() {
  /*
    should the program name be included ?
    im ripping it out coz easier to implement commands
    args[0] = first command not "main.exe" / "super".
    might change
  */
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
      /* i think this is a dumb way to do it but it works */
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

/*
  should prob move to init.go coz only used in "init" for now.
  but keeping it here as "commit" and "add" might need to check if repo exists or smth.
  move this shit if they dont.
  or add to a "utils.go" or smth.
*/
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}
