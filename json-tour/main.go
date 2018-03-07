package main

import "fmt"
import "encoding/json"

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	fmt.Println("Don't panic, we're testing json stuff...")

	// first part: marshalling
	m := Message{"Arthur", "my house", 1294706395881547000}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("serialized: ", string(b))

	// second part: unmarshal
	var m2 Message
	b2 := []byte(`{"Name": "Bob", "Food":"Pizza"}`)
	err = json.Unmarshal(b2, &m2)

	fmt.Println("input string: ", string(b2))
	fmt.Println("desserialized: ", m2)
	fmt.Println("(the deserialized version should not contain the" +
		" field \"Food\" and the time is 0)")
}
