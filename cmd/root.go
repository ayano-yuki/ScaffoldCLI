// cmd/root.go
package cmd

import (
	"fmt"
	"os"
)

func Root() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scaffold <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		if err := InitCommand(); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
