package main

import (
	"fmt"
	"mre/mypackage1"
	"strings"
)

func main() {
	if !strings.Contains(mypackage1.MyVar.Name, "mypackage1") {
		fmt.Printf("caller file should contain 'mypackage1', but was '%s'", mypackage1.MyVar.Name)
	}
}
