package main

import (
	"fmt"
	"os"
	"path"

	"github.com/vektah/dataloaden/pkg/generator"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: name keyType valueType [dir]")
		fmt.Println(" example:")
		fmt.Println(" dataloaden 'UserLoader int []*github.com/my/package.User'")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if len(os.Args) == 5 {
		if path.IsAbs(os.Args[4]) {
			wd = os.Args[4]
		} else {
			wd = path.Join(wd, os.Args[4])
		}
	}

	if err := generator.Generate(os.Args[1], os.Args[2], os.Args[3], wd); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
