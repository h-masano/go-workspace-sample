package main

import (
	"fmt"
	"log"

	"runtime/debug"
	"github.com/samber/lo"
)

func main() {
	lo.ToPtr("module-f")

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}
	fmt.Println("----------------------module-f----------------------")
	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}
}