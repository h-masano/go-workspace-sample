package main

import (
	"fmt"
	"log"
	// modd "module-d"
	"runtime/debug"
	"github.com/samber/lo"
)

func main() {
	// d := modd.ModD()
	// fmt.Println("run modc, import", d)
	// lo.ToPtr(d)
	lo.ToPtr("module-c")

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}
	fmt.Println("----------------------module-c----------------------")
	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}
}
