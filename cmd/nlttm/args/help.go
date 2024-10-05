package _args

import (
	"fmt"
	"os"
)

func Help(missingArg bool, arg string) {
	if missingArg {
		fmt.Println(true)
	} else {
		fmt.Println(false, arg)
		os.Exit(1)
	}
}
