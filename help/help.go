package help

import (
  "fmt"
)

func OutputHelpCommand() {
  const help_message string = 
  `
     sssssssssssssss uuuuuuuu     uuuuuuuuppppppppppppppppp   eeeeeeeeeeeeeeeeeeeeeerrrrrrrrrrrrrrrrr   
 ss:::::::::::::::su::::::u     u::::::up::::::::::::::::p  e::::::::::::::::::::er::::::::::::::::r  
s:::::ssssss::::::su::::::u     u::::::up::::::pppppp:::::p e::::::::::::::::::::er::::::rrrrrr:::::r 
s:::::s     sssssssuu:::::u     u:::::uupp:::::p     p:::::pee::::::eeeeeeeee::::err:::::r     r:::::r
s:::::s             u:::::u     u:::::u   p::::p     p:::::p  e:::::e       eeeeee  r::::r     r:::::r
s:::::s             u:::::d     d:::::u   p::::p     p:::::p  e:::::e               r::::r     r:::::r
 s::::ssss          u:::::d     d:::::u   p::::pppppp:::::p   e::::::eeeeeeeeee     r::::rrrrrr:::::r 
  ss::::::sssss     u:::::d     d:::::u   p:::::::::::::pp    e:::::::::::::::e     r:::::::::::::rr  
    sss::::::::ss   u:::::d     d:::::u   p::::ppppppppp      e:::::::::::::::e     r::::rrrrrr:::::r 
       ssssss::::s  u:::::d     d:::::u   p::::p              e::::::eeeeeeeeee     r::::r     r:::::r
            s:::::s u:::::d     d:::::u   p::::p              e:::::e               r::::r     r:::::r
            s:::::s u::::::u   u::::::u   p::::p              e:::::e       eeeeee  r::::r     r:::::r
sssssss     s:::::s u:::::::uuu:::::::u pp::::::pp          ee::::::eeeeeeee:::::err:::::r     r:::::r
s::::::ssssss:::::s  uu:::::::::::::uu  p::::::::p          e::::::::::::::::::::er::::::r     r:::::r
s:::::::::::::::ss     uu:::::::::uu    p::::::::p          e::::::::::::::::::::er::::::r     r:::::r
 sssssssssssssss         uuuuuuuuu      pppppppppp          eeeeeeeeeeeeeeeeeeeeeerrrrrrrr     rrrrrrr


 welcome to super, the version control manager that is not so super (for now)
 super is meant to be a lightweight, easy-to-use and scalable alternative to git.
 
 basic commands:
 super init     - initialises a repository.
 super add      - adds current changes to commit.
 super commit   - commits the current changes to repository.


  `

  fmt.Println(help_message)
}
