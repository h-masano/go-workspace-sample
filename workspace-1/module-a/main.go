package main

import (
	"fmt"
	"log"

	// "github.com/h-masano/go-workspace-sample/workspace-1/module-b"
	"runtime/debug"
	"github.com/samber/lo"
)

func main() {
	// b := modb.ModB()
	// fmt.Println("run moda, import ", b)
	lo.ToPtr("module-a")

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}
	fmt.Println("----------------------module-a----------------------")
	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}
}