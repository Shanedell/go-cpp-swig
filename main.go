package main

import (
	"fmt"

	"github.com/Shanedell/go-cpp-swig/gocpp"
)

func main() {
	var math = gocpp.NewMath()
	fmt.Println(math.Sum(1, 2))
}
