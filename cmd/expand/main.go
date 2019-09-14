package main

import (
	"fmt"
	"log"

	"github.com/Raita876/expand"
)

func main() {
	qpc, err := expand.Parse("expand.yaml")
	if err != nil {
		log.Fatal(err)
	}

	urlArray, err := qpc.Create()
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range urlArray {
		fmt.Printf("%s\n", url)
	}

}
