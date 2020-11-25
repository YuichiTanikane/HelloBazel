package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("---start---")

	fmt.Printf("new uuid is [%s]\n", makeUUID())

	fmt.Println("---end---")
}

func makeUUID() string {
	return uuid.New().String()
}
