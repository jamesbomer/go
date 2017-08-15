package main

import (
	"fmt"

	"github.com/jamesbomer/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse(stringutil.Reverse("hello world!\n")))
}
