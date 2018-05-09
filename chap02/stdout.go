package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println()
	os.Stdout.Write([]byte("os.Stdout example\n"))
}