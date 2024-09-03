package modd

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/samber/lo"
)

func ModD() string {
	str := "modd"
	lo.ToPtr(str)

	// show module versions in mod-d
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return "fail in modd"
	}

	fmt.Println("----------------------module-d----------------------")
	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}
	return str
}
