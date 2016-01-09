package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(
		//Shamsher Ansari
		strings.Trim("  Shamsher Ansari  ", " "),
		// true
		strings.Contains("Shamsher Ansari", "Shams"),
		//SHAMSHER
		strings.ToUpper("shamsher"),
		//shamsher
		strings.ToLower("SHAMSHER"),
		//4
		strings.Index("Shamsher", "sher"),
		//Shh... Shh... Shh...
		strings.Repeat("Shh... ", 3),
	)

}
