package main

import (
	"fmt"

	color "github.com/fatih/color"
)

func main() {

	actor := color.New(color.FgCyan).SprintFunc()
	quote := color.New(color.FgWhite).SprintFunc()

	fmt.Printf("%s %s\n", actor("Frank:"), quote("How can you trust a man that wears both a belt and suspenders? Man can't even trust his own pants."))

}
