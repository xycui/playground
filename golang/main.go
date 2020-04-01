package main

import (
	"fmt"
	"os"

	"github.com/xycui/playground/container"
)

func main() {
	builder := container.Level0Builder
	cmd := builder.Build(os.Args[1:])
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
