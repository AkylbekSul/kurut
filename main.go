package main

import (
	"os"

	"github.com/rasulov-emirlan/kurut/runtime"
)

func main() {
	file, err := os.Open("testing/example.kut")
	if err != nil {
		panic(err)
	}

	err = runtime.Run(file)
	if err != nil {
		panic(err)
	}
}