package help

import (
  "fmt"
)


/* the text is really weirdly formatted but idgaf */
func OutputHelpCommand() {
  const help_message string = 
`
  ___ _   _ _ __   ___ _ __ 
 / __| | | | '_ \ / _ \ '__|
 \__ \ |_| | |_) |  __/ |   
 |___/\__,_| .__/ \___|_|   
           | |              
           |_|              

 welcome to super, the version control manager that is (not so) super.
 super is meant to be a lightweight, easy-to-use and scalable alternative to git.
 key words here: 'meant to be'.
 
 basic commands:
 super init     - initialises a repository.
 super add      - adds current changes to commit.
 super commit   - commits the current changes to repository.`

  /* should this return the message instead of printing ? */
  fmt.Println(help_message)
}
