package main

import (
	"flag"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// "-i {{UUIDを発行したい数}}"" でその個数だけ発行できる。オプションなしで5個発行する。
var countOpt = flag.Int("i", 5, "only \"i\" option")

func main() {

	flag.Parse()

	for i := 0; i < *countOpt; i++ {
		fmt.Printf("%v\n", uuid.NewV4().String())
	}

}
