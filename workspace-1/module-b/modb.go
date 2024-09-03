package modb

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/samber/lo"
)

func ModB() string {
	str := "modb"
	lo.ToPtr(str)

	// show module versions in mod-d
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return "fail in modd"
	}

	fmt.Println("----------------------module-b----------------------")
	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}
	return str
}