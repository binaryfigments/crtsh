package main

import (
	"encoding/json"
	"fmt"

	"github.com/binaryfigments/crtsh"
)

func main() {
	data := crtsh.Get("networking4all.com", 5)

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}
