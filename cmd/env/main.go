package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i, pair := range os.Environ() {
		kv := strings.Split(pair, "=")
		fmt.Println(fmt.Printf("%d-%s-%s", i, kv[0], kv[1]))

	}
}
