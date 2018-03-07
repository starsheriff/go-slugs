package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Token struct {
	Key string
	Age int
}

func main() {
	argsWithProg := os.Args
	args := os.Args[1:]

	fmt.Println("All args: ", argsWithProg)
	fmt.Println("args: ", args)

	var b Token
	err := json.Unmarshal([]byte(args[0]), &b)
	if err != nil {
		log.Fatal("could not parse input: ", err, b)
	}

	fmt.Println("parsed: ", b)
}
