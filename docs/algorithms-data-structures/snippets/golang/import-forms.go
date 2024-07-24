package main

import (
	"fmt"

	/*
	   In fact, the full form of an import declaration is
	    ```
	    import importname "path/to/package"
	    ```
	   Where importname is optional, its default value is the name
	   (not the folder name) of the imported package.


	   These import declarations are equivalent to the following ones:
	    ```
	    import fmt "fmt"        // <=> import "fmt"
	    import rand "math/rand" // <=> import "math/rand"
	    import time "time"      // <=> import "time"
	    ```

	*/

	/*
		The `importname` in the full form import declaration can be a dot (`.`).
		Such imports are called dot imports.
		To use the exported elements in the packages being dot imported, the prefix
		part in qualified identifiers must be omitted.
	*/
	// generally, dot imports reduce code readability, so they are not
	// recommended to be used in formal projects.
	. "fmt"
	. "time"

	/*
		The `importname` in the full form import declaration can be the blank
		identifier (`_`).
		Such imports are called anonymous imports (some articles elsewhere also
		call them blank imports).
	*/
	format "fmt"  // okay: it is used once below
	_ "math/rand" // okay: it is not required to be used, due to the blank identifier
)

func main() {
	Println("Current time:", Now()) // due to the dot import, the prefix part is omitted
	format.Println()                // use the imported "fmt" package
	fmt.Println("vim-go")
}
