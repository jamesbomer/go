package main

import (
	"fmt"

	"github.com/jamesbomer/go/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse(stringutil.Reverse("hello world!\n")))
}
