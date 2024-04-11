package help

import (
  "fmt"
)

func OutputHelpCommand() {
  const help_message string = 
  `                                                                                                      
   SSSSSSSSSSSSSSS UUUUUUUU     UUUUUUUUPPPPPPPPPPPPPPPPP   EEEEEEEEEEEEEEEEEEEEEERRRRRRRRRRRRRRRRR   
 SS:::::::::::::::SU::::::U     U::::::UP::::::::::::::::P  E::::::::::::::::::::ER::::::::::::::::R  
S:::::SSSSSS::::::SU::::::U     U::::::UP::::::PPPPPP:::::P E::::::::::::::::::::ER::::::RRRRRR:::::R 
S:::::S     SSSSSSSUU:::::U     U:::::UUPP:::::P     P:::::PEE::::::EEEEEEEEE::::ERR:::::R     R:::::R
S:::::S             U:::::U     U:::::U   P::::P     P:::::P  E:::::E       EEEEEE  R::::R     R:::::R
S:::::S             U:::::D     D:::::U   P::::P     P:::::P  E:::::E               R::::R     R:::::R
 S::::SSSS          U:::::D     D:::::U   P::::PPPPPP:::::P   E::::::EEEEEEEEEE     R::::RRRRRR:::::R 
  SS::::::SSSSS     U:::::D     D:::::U   P:::::::::::::PP    E:::::::::::::::E     R:::::::::::::RR  
    SSS::::::::SS   U:::::D     D:::::U   P::::PPPPPPPPP      E:::::::::::::::E     R::::RRRRRR:::::R 
       SSSSSS::::S  U:::::D     D:::::U   P::::P              E::::::EEEEEEEEEE     R::::R     R:::::R
            S:::::S U:::::D     D:::::U   P::::P              E:::::E               R::::R     R:::::R
            S:::::S U::::::U   U::::::U   P::::P              E:::::E       EEEEEE  R::::R     R:::::R
SSSSSSS     S:::::S U:::::::UUU:::::::U PP::::::PP          EE::::::EEEEEEEE:::::ERR:::::R     R:::::R
S::::::SSSSSS:::::S  UU:::::::::::::UU  P::::::::P          E::::::::::::::::::::ER::::::R     R:::::R
S:::::::::::::::SS     UU:::::::::UU    P::::::::P          E::::::::::::::::::::ER::::::R     R:::::R
 SSSSSSSSSSSSSSS         UUUUUUUUU      PPPPPPPPPP          EEEEEEEEEEEEEEEEEEEEEERRRRRRRR     RRRRRRR

 Welcome to SUPER, the version control manager that is not so super (for now)
 Super is meant to be a lightweight, easy-to-use and scalable alternative to Git.
 
 Basic Commands:
 super init     - Initialises a repository.
 super add      - Adds current changes to commit.
 super commit   - Commits the current changes to repository.
  `

  fmt.Println(help_message)
}
