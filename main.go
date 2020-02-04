// A self-reproducing program that satisfies gofmt.
// Brad Tilley <btilley@gatech.edu>

package main

import "fmt"

func main() {
	s := "// A self-reproducing program that satisfies gofmt.\n// Brad Tilley <btilley@gatech.edu>\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\ts := %q\n\tfmt.Printf(s, s)\n}\n"
	fmt.Printf(s, s)
}
