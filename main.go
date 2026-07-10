package main

import (
	"fmt"

	"github.com/nerdnavya/mydb/kv"
)

func main() {
	err := kv.SaveData("test.db", []byte("hello database!"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Success! Check test.db")
}
