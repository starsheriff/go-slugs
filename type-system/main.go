package main

import "fmt"

type wannabeenum uint8

const (
	zero wannabeenum = iota
	one  wannabeenum = iota
)

func takeWannaBe(a wannabeenum) {
	fmt.Println("a is: ", a)
}

func main() {
	fmt.Println("Let's call `takeWannaBe` with a wannabeenum constant.")
	takeWannaBe(zero)
	takeWannaBe(one)

	fmt.Println("Can we call `takeWannaBe` with its _underlying type_?")
	takeWannaBe(1)

	fmt.Println("Yes we can, this is weird. Go is supposed to be strongly typed.")

	fmt.Println("Ok, but these calls do not compile:")
	fmt.Println("takeWannaBe(uint64(1))")
	//takeWannaBe(uint64(1))
	fmt.Println("takeWannaBe(uint8(1))")
	//takeWannaBe(uint8(2))

	fmt.Println("This means that the in the first call we have a type inference happening")
	fmt.Println("What about `takeWannaBe(10)` then? (10 is not defined in const)")
	takeWannaBe(10)
	fmt.Println("It works, of course. The underlying type is a uint8")
}
