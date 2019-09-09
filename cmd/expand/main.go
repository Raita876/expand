package main

import (
	"fmt"
	"os"

	"github.com/Raita876/expand"
)

func main() {
	qpc, err := expand.Parse("expand.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(qpc)
}
