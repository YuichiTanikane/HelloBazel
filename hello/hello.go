package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("---start---")

	fmt.Printf("test id = [%s]\n", makeUUID())

	fmt.Println("---end---")
}

func makeUUID() string {
	return uuid.New().String()
}
