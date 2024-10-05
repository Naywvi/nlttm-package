package _args

import (
	"fmt"
	"os"
)

func Help(missingArg bool, arg string) {
	if missingArg {
		fmt.Println("help func true")
	} else {
		fmt.Println("help func false", arg)
		os.Exit(1)
	}
}
